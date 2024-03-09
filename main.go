package main

import (
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var stopReason string
var port int

func init() {
	flag.IntVar(&port, "port", 3000, "Port to run the server on")
	flag.Parse()
}

func main() {
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	stopReason = "stop"
	server := gin.Default()
	server.Use(CORS())
	server.POST("/v1/chat/completions", func(c *gin.Context) {
		var chatRequest ChatRequest
		if err := c.ShouldBindJSON(&chatRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		prompt := "This is a mock server."
		if len(chatRequest.Messages) != 0 {
			prompt = chatRequest.Messages[len(chatRequest.Messages)-1].Content
		}
		response := prompt2response(prompt)

		if chatRequest.Stream {
			setEventStreamHeaders(c)
			dataChan := make(chan string)
			stopChan := make(chan bool)
			streamResponse := ChatCompletionsStreamResponse{
				Id:      "chatcmpl-iwillalwaysloveyou",
				Object:  "chat.completion.chunk",
				Created: 1689411338,
				Model:   "gpt-3.5-turbo",
			}
			streamResponseChoice := ChatCompletionsStreamResponseChoice{}
			go func() {
				for i, s := range response {
					streamResponseChoice.Delta.Content = string(s)
					if i == len(response)-1 {
						streamResponseChoice.FinishReason = &stopReason
					}
					streamResponse.Choices = []ChatCompletionsStreamResponseChoice{streamResponseChoice}
					jsonStr, _ := json.Marshal(streamResponse)
					dataChan <- string(jsonStr)
				}
				stopChan <- true
			}()

			c.Stream(func(w io.Writer) bool {
				select {
				case data := <-dataChan:
					c.Render(-1, CustomEvent{Data: "data: " + data})
					return true
				case <-stopChan:
					c.Render(-1, CustomEvent{Data: "data: [DONE]"})
					return false
				}
			})
		} else {
			c.JSON(http.StatusOK, Completion{
				Id:      "chatcmpl-7f8Qxn9XkoGsVcl0RVGltZpPeqMAG",
				Object:  "chat.completion",
				Created: time.Now().Unix(),
				Model:   "gpt-3.5-turbo",
				Choices: []Choice{
					{
						Index: 0,
						Message: Message{
							Role:    "assistant",
							Content: prompt,
						},
						FinishReason: "length",
					},
				},
				Usage: Usage{
					PromptTokens:     9,
					CompletionTokens: 1,
					TotalTokens:      10,
				},
			})
		}
	})

	log.Printf("Starting server on port %d", port)
	log.Fatal(server.Run(":" + strconv.Itoa(port)))
}
