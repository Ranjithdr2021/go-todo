package Routes

import (
	"go-todo-app/Controllers"
	middleware "go-todo-app/middlewares"

	"github.com/gin-gonic/gin"
)

var encryptedString string

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.Use(middleware.DecryptRequest())
	{
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
		v1.POST("/token", Controllers.Login)
		v1.POST("/user/register", Controllers.RegisterUser)
		secured := v1.Group("/secured")
		{
			secured.GET("/ping", Controllers.Ping)
		}
	}
	encrypt := r.Group("/data")
	{
		encrypt.POST("encrypt", Controllers.EncryptDataHandler)
		encrypt.POST("decrypt", Controllers.DecryptDataHandler)
	}
	return r
}
