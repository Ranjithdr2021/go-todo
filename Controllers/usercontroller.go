package Controllers

import (
	"encoding/json"
	"errors"
	"go-todo-app/Config"
	"go-todo-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	decryptedData, exists := context.Get("decryptedText")
	if !exists {
		context.AbortWithError(http.StatusBadRequest, errors.New("decrypted data not found"))
		return
	}
	json.Unmarshal(decryptedData.([]byte), &user)
	db := Config.Database.ConnectToDB()
	defer db.Close()
	_, err := db.Query("insert into users(Name, Username, Email, Password) values(?,?,?,?)", user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, AESEncrypt("Success........", []byte(context.Request.Header.Get("x-key")), context.Request.Header.Get("x-iv")))
}
