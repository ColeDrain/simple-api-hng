package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

    router.GET("/anon", getme)

	port := os.Getenv("PORT")
    router.Run("0.0.0.0:"+port)
}

// getme responds with me.
func getme(c *gin.Context) {
	// Add CORS headers
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

    c.IndentedJSON(http.StatusOK, users[0])
}
