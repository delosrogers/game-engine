package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Hub maintains the set of active rooms and helps add new clients to them
type Hub struct {
	rooms map[uuid.UUID]Room
}

func newHub() *Hub {
	return &Hub{
		rooms: make(map[uuid.UUID]Room),
	}
}

// JoinRoom allows a client to connect to a room
func (h *Hub) addClientToRoom(roomID uuid.UUID, c *gin.Context) {
	room := h.rooms[roomID]
	serveWs(&room, c.Writer, c.Request)
}

// createRoom create a new room and starts it with a max capacity, defaults to 2
func (h *Hub) createRoom(args ...int) (uuid.UUID, error) {
	var maxClients int
	if len(args) > 0 {
		maxClients = args[0]
	} else {
		maxClients = 2
	}
	roomId, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}
	room := newRoom(maxClients)
	h.rooms[roomId] = *room
	go room.run()
	return roomId, nil
}


