package hpsdk

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyperits/gosuite/logger"
)

type OpenApiMiddleware struct {
	ApiKeys map[string]string
	SignTtl int
}

func NewOpenApiMiddleware(apiKeys map[string]string, signTtl int) *OpenApiMiddleware {
	return &OpenApiMiddleware{
		ApiKeys: apiKeys,
		SignTtl: signTtl,
	}
}

func (mw *OpenApiMiddleware) GetSecretByAccessKey(accessKey string) (string, error) {
	if apiKey, ok := mw.ApiKeys[accessKey]; ok {
		return apiKey, nil
	}
	return "", errors.New("[hpsdk] no such accessKey")
}

func (mw *OpenApiMiddleware) OpenApiAuth(c *gin.Context) {

	accessKey := c.Query(OpenApiAccessKey)
	timestamp := c.Query(OpenApiTimestamp)
	sign := c.Query(OpenApiSign)

	iTimestamp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		logger.Errorf("[hpsdk] %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	duration := time.Since(time.Unix(iTimestamp, 0))
	if duration > time.Duration(mw.SignTtl)*time.Minute {
		logger.Errorf("[hpsdk] [%v, %v, %v] timestamp is expired", accessKey, timestamp, sign)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	secret, err := mw.GetSecretByAccessKey(accessKey)
	if err != nil {
		logger.Errorf("[hpsdk] %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	serverSign := GenerateHMACSHA256Digest(accessKey, secret, iTimestamp)

	if sign != serverSign {
		logger.Errorf("[hpsdk] [%v, %v]", sign, serverSign)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
