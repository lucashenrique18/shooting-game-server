package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucashenrique18/shooting-game-server/internal/infra/adapters"
	"github.com/lucashenrique18/shooting-game-server/internal/infra/modules"
	"github.com/lucashenrique18/shooting-game-server/internal/infra/websocket"
)

func main() {
	println("Starting game server")
	controllersModules, _, _ := modules.InitializeModules(uint16(60))
	createWebSocket(controllersModules)
	createHttpRoutes(controllersModules)
}

func createWebSocket(controllers modules.ControllersModules) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(w, r, controllers.PlayerEventController)
	})
	go http.ListenAndServe(":8080", nil)
	fmt.Println("WebSocket rodando em ws://localhost:8080/ws")
}

func createHttpRoutes(controllers modules.ControllersModules) {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.POST("/create-match", adapters.AdaptRoute(controllers.CreateMatchController))
	r.POST("/join-match", adapters.AdaptRoute(controllers.JoinMatchController))
	r.GET("/matches", adapters.AdaptRoute(controllers.GetAllPossibleMatchesController))
	r.Run(":8081")
	fmt.Println("HTTP com Gin rodando em http://localhost:8081")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
