package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `json:"username" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
}

var logins = []login{

}

func main() {
	router := gin.Default()
	router.POST("/sign-up", loginUser)
	router.POST("/tweets", )
	router.GET("/tweets", )

	router.Run("localhost:4000")
}


func loginUser(user *gin.Context) {
	var newLogin login

	if err := user.BindJSON(&newLogin); err != nil {
		user.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	user.BindJSON(&newLogin)
	logins = append(logins, newLogin)
	user.IndentedJSON(http.StatusCreated, "User created")
}