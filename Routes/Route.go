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
	add := v1.Group("/add")
	add.Use(middleware.DecryptRequest())
	{
		add.POST("todo", Controllers.CreateATodo)
		add.POST("/user/register", Controllers.RegisterUser)
		add.PUT("todo/:id", Controllers.UpdateATodo)
		v1.GET("todo", Controllers.GetTodos)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
		v1.POST("/token", Controllers.Login)

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
