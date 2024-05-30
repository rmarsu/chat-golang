package handler

import (
	chat "chat"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	chat.Service
	hub *Hub
}

func NewHandler(s chat.Service, h *Hub) *Handler {
	return &Handler{
		Service: s,
		hub:     h,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.CreateUser)
		auth.POST("/sign-in", h.Login)
	}
	websocket := router.Group("/ws")
	{
		websocket.POST("/createRoom", h.CreateRoom)
		websocket.GET("/joinRoom/:roomId", h.JoinRoom)
		websocket.GET("/joinRooms", h.GetRooms)
		websocket.GET("/getClients/:roomId", h.GetClients)
	}
	return router
}
