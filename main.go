package main

import (
	"net/http"
	"os"

	"time"

	"github.com/gin-gonic/gin"

	cors "github.com/itsjamie/gin-cors"
)

// user represents data about user.
type user struct {
    SlackUsername     string  `json:"slackUsername"`
    Backend  bool  `json:"backend"`
    Age int  `json:"age"`
    Bio  string `json:"bio"`
}

// albums slice to seed record album data.
var users = []user{
    { SlackUsername: "anonymous", Backend: true, Age: 23, Bio: "I am everywhere, more like the Davinci of Tech."},
}

func main() {
    router := gin.Default()

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: false,
		ValidateHeaders: false,
	}))

    router.GET("/anon", getme)

	port := os.Getenv("PORT")
    router.Run("localhost:"+port)
}

// getme responds with me.
func getme(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users[0])
}
