package Controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go-todo-app/Config"
	"go-todo-app/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	// db := Config.ConnectToDB()
	db := Config.Database.ConnectToDB()
	defer db.Close()
	row, err := db.Query("SELECT * FROM todo")
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}
	for row.Next() {
		var todo models.Todo
		if err := row.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
			fmt.Fprint(c.Writer, err)
			return
		}
		todos = append(todos, todo)
	}
	data, _ := json.Marshal(todos)
	fmt.Println(string(data))
	c.JSON(http.StatusOK, AESEncrypt(string(data), []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	decryptedData, exists := c.Get("decryptedText")
	if !exists {
		c.AbortWithError(http.StatusBadRequest, errors.New("decrypted data not found"))
		return
	}
	json.Unmarshal(decryptedData.([]byte), &todo)
	db := Config.Database.ConnectToDB()
	defer db.Close()
	_, err := db.Query("insert into todo(Title, Description) values(?,?)", todo.Title, todo.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, AESEncrypt("Todo created Successfully.....", []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))
}

func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	db := Config.Database.ConnectToDB()
	defer db.Close()
	row, err := db.Query("SELECT * FROM todo where ID=?", id)
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}
	for row.Next() {
		if err := row.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
			fmt.Fprint(c.Writer, err)
			return
		}
	}
	data, _ := json.Marshal(todo)
	fmt.Println(data)
	c.JSON(http.StatusOK, AESEncrypt(string(data), []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))
}

func UpdateATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	decryptedData, exists := c.Get("decryptedText")
	if !exists {
		c.AbortWithError(http.StatusBadRequest, errors.New("decrypted data not found"))
		return
	}
	json.Unmarshal(decryptedData.([]byte), &todo)
	db := Config.Database.ConnectToDB()
	defer db.Close()
	_, err := db.Exec("update todo set Title=?, Description=? where ID=?", todo.Title, todo.Description, id)
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, AESEncrypt("Updated Successfully.......", []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))
}

func DeleteATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	db := Config.Database.ConnectToDB()
	defer db.Close()
	_, err := db.Exec("DELETE from todo where ID=?", id)
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, AESEncrypt("Record deleted Succesfully.......", []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))
}
