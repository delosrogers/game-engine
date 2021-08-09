package core

import (
	"github.com/delosrogers/game-engine/pkg/gamework"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Hub maintains the set of active rooms and helps add new clients to tshouhem
type Hub struct {
	rooms map[uuid.UUID]Room
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[uuid.UUID]Room),
	}
}

// addClientToRoom allows a client to connect to a room
func (h *Hub) AddClientToRoom(roomID uuid.UUID, c *gin.Context) {
	room := h.rooms[roomID]
	serveWs(&room, c.Writer, c.Request)
}

// createRoom create a new room and starts it with a max capacity, defaults to 2
func (h *Hub) CreateRoom(game *gamework.Game) (uuid.UUID, error) {
	roomId, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}
	room := newRoom(game)
	h.rooms[roomId] = *room
	go room.run()
	return roomId, nil
}


