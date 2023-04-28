package Controllers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EncryptDataHandler(ctx *gin.Context) {
	var requestBody interface{}
	ctx.ShouldBindJSON(&requestBody)
	plainText, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(requestBody)
	encryptedData := AESEncrypt(string(plainText), []byte(ctx.Request.Header.Get("x-key")), ctx.Request.Header.Get("x-iv"))
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
	ctx.JSON(http.StatusOK, encryptedString)
}

func DecryptDataHandler(ctx *gin.Context) {
	var requestBody map[string]string
	ctx.ShouldBindJSON(&requestBody)
	fmt.Println(requestBody, requestBody["Encrypted-data"])
	encryptedData, _ := base64.StdEncoding.DecodeString(requestBody["Encrypted-data"])
	decryptedText := AESDecrypt(encryptedData, []byte(ctx.Request.Header.Get("x-key")), ctx.Request.Header.Get("x-iv"))
	fmt.Println("\n decrypted data:", string(decryptedText))
	ctx.JSON(http.StatusOK, string(decryptedText))
}

func AESEncrypt(src string, key []byte, IV string) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(IV))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return (crypted)

}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AESDecrypt(crypt []byte, key []byte, IV string) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(IV))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)
	return PKCS5Trimming(decrypted)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
