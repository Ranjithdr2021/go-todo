package middleware

import (
	"encoding/base64"
	"fmt"
	"go-todo-app/Controllers"

	"github.com/gin-gonic/gin"
)

func DecryptRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody map[string]string
		ctx.ShouldBindJSON(&requestBody)
		fmt.Println(requestBody, requestBody["Encrypted-data"])
		encryptedData, _ := base64.StdEncoding.DecodeString(requestBody["Encrypted-data"])
		decryptedText := Controllers.AESDecrypt(encryptedData, []byte(ctx.Request.Header.Get("x-key")), ctx.Request.Header.Get("x-iv"))
		fmt.Println("\n decrypted data:", string(decryptedText))
		ctx.Set("decryptedText", decryptedText)
		ctx.Next()
	}
}

// func EncryptResponse() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		ctx.Next()
// 		var responseData []byte
// 		var err error
// 		switch v := ctx.Writer.(type) {
// 		case *gzip.Writer:
// 			// If the response is compressed, get the compressed bytes
// 			responseData = v.Writer.(*bytes.Buffer).Bytes()
// 		default:
// 			responseData = ctx.Writer.(*bytes.Buffer).Bytes()
// 		}
// 	}
// }
