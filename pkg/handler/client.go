package handler

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomid"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomid"`
	Username string `json:"username"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func () {
		hub.UnRegister <- c
          c.Conn.Close()
	}()
	
	for {	
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Printf("ошибочка: %v", err)
		}
		break
	}
	msg := &Message{
		Content: string(m),
		RoomID: c.RoomID,
		Username: c.Username,
	} 
	hub.Broadcast <- msg
	}
}


