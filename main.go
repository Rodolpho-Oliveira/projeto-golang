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

	if err := user.ShouldBindJSON(&newLogin); err != nil {
		user.JSON(http.StatusBadRequest, "Missing information")
		return
	}

	user.ShouldBindJSON(&newLogin)
	logins = append(logins, newLogin)
	user.IndentedJSON(http.StatusAccepted, "User created")
}