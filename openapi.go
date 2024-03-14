package hpsdk

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperits/gosuite/logger"
)

const (
	OpenApiAccessKey = "accessKey"
	OpenApiTimestamp = "timestamp"
	OpenApiSign      = "sign"
)

type OpenApiArgs struct {
	AccessKey string
	Timestamp int64
	Sign      string
}

func GetOpenApiArgs(c *gin.Context) (*OpenApiArgs, error) {
	accessKey := c.Query(OpenApiAccessKey)
	timestamp := c.Query(OpenApiTimestamp)
	sign := c.Query(OpenApiSign)

	iTimestamp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		logger.Errorf("[hpsdk] failed to get int timestamp from header")
		return nil, err
	}

	return &OpenApiArgs{
		AccessKey: accessKey,
		Timestamp: iTimestamp,
		Sign:      sign,
	}, nil
}
