package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `json:"username" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
}

type tweet struct {
	Username string `json:"username" binding:"required"`
	Tweet string `json:"tweet" binding:"required"`
}

var logins = []login{}
var tweets = []tweet{
	{Username: "eae", Tweet: "eaeae"},
}

func main() {
	router := gin.Default()
	router.POST("/sign-up", loginUser)
	router.POST("/tweets", createTweet)
	router.GET("/tweets", getAllTweets)

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
	user.IndentedJSON(http.StatusCreated, "User created")
}

func createTweet(reqTweets *gin.Context) {
	var newTweet tweet

	if err := reqTweets.ShouldBindJSON(&newTweet); err != nil {
		reqTweets.JSON(http.StatusBadRequest, "Missing information")
		return
	}

	reqTweets.ShouldBindJSON(&newTweet)
	if(len(tweets) == 10) {
		tweets = tweets[1:]
	}
	tweets = append(tweets, newTweet)
	reqTweets.IndentedJSON(http.StatusCreated, "Tweet received")
}

func getAllTweets(getTweets *gin.Context) {
	getTweets.IndentedJSON(http.StatusAccepted, tweets)
}