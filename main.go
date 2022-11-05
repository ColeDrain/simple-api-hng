package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
)

// user represents data about user.
type user struct {
    SlackUsername     string  `json:"slackUsername"`
    Backend  bool  `json:"backend"`
    Age int  `json:"age"`
    Bio  string `json:"bio"`
}

type RequestOperation struct {
	Op string `json:"operation_type,omitempty"`
	X             int    `json:"x,omitempty"`
	Y             int    `json:"y,omitempty"`
}

type ResponseOperation struct {
	SlackUsername string `json:"slackUsername,omitempty"`
	Result        int    `json:"result,omitempty"`
	OperationType string `json:"operation_type,omitempty"`
	Error         string `json:"error,omitempty"`
}

var users = []user{
    { SlackUsername: "anonymous", Backend: true, Age: 23, Bio: "I am everywhere, more like the Davinci of Tech."},
}

func main() {
    router := gin.Default()

    router.GET("/anon", getme)
    router.POST("/calculate", calculate)

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


func calculate(c *gin.Context) {
    var newOp RequestOperation

    resp := new(ResponseOperation)
    resp.SlackUsername = "SonOfVinci"

    // Call BindJSON to bind the received JSON to newOperation
    if err := c.BindJSON(&newOp); err != nil {
        c.IndentedJSON(http.StatusBadRequest, err)
    }

    switch newOp.Op {
    case "addition":
        resp.Result = newOp.X + newOp.Y
    case "subtraction":
        resp.Result = newOp.X - newOp.Y
    case "multiplication":
        resp.Result = newOp.X * newOp.Y
    default:
        fmt.Println(gptCalculate(newOp.Op))

        res, err := gptCalculate(newOp.Op)
        if err != nil {
            resp.Error = "Bad Input 🤔"
        } else {
            resp.Result = res
        }
    }

    resp.OperationType = newOp.Op

    c.IndentedJSON(http.StatusOK, resp)
}

func gptCalculate(prompt string) (int, error){
    c := gogpt.NewClient("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")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model: "text-davinci-002",
        Prompt: prompt,
        Temperature: 0.3,
        MaxTokens: 60,
        TopP: 1,
        FrequencyPenalty: 0.8,
        PresencePenalty: 0,
	}

    resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return -1, err
	}

    // Result from OpenAI
    result := resp.Choices[0].Text

    // Get only the number if possible
    s := strings.Split(result, "=")
    slast := s[len(s)-1]
    strim := strings.TrimSpace(slast)
    i, err := strconv.Atoi(strim)
    if err != nil {
        return -1, err
    }
    return i, nil
}