package hpsdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// GenerateHMACSHA256Digest 用于生成摘要
func GenerateHMACSHA256Digest(appID, secretKey string, timestamp int64) string {

	message := fmt.Sprintf("%s:%d", appID, timestamp)

	// 使用 HMAC-SHA256 算法生成摘要
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(message))
	digest := hex.EncodeToString(h.Sum(nil))

	return digest
}
