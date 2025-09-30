package adapters

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	controllers "github.com/lucashenrique18/shooting-game-server/internal/presentation"
)

func AdaptRoute(controller controllers.ControllerInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := make(map[string]interface{})
		contentType := c.GetHeader("Content-Type")

		if c.Request.Method != "GET" {
			switch {
			case strings.Contains(contentType, "application/json"):
				if err := c.ShouldBindJSON(&request); err != nil {
					if c.Request.ContentLength != 0 {
						c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
						return
					}
				}
			case strings.Contains(contentType, "application/x-www-form-urlencoded"):
				// Handle regular form data
				form := c.Request.PostForm
				for key, values := range form {
					if len(values) > 0 {
						request[key] = values[0]
					}
				}
			}
		}

		for key, values := range c.Request.URL.Query() {
			if len(values) > 0 {
				request[key] = values[0]
			}
		}

		for _, param := range c.Params {
			request[param.Key] = param.Value
		}

		statusCode, response := controller.Handle(request)
		if statusCode > 299 {
			fmt.Println("Error handling request", map[string]interface{}{"statusCode": statusCode, "response": response})
		}
		c.JSON(statusCode, response)
	}
}
