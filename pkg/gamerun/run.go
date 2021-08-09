package gamerun

import (
	"fmt"
	"github.com/delosrogers/game-engine/internal/core"
	"github.com/delosrogers/game-engine/pkg/gamework"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/url"
	"os"
)

func Run(game gamework.Game, addr ...string) {
	hub := core.NewHub()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		roomId, err := hub.CreateRoom(game50)
		if err != nil {
			log.Fatal("err creating room: ", err)
		}
		c.Redirect(307, url.PathEscape(roomId.String()))
	})
	router.GET("/:id", serveHome)
	//router.GET("/game/:id"), handleGame)
	router.GET("/:id/ws", func(c *gin.Context) {
		roomIdString, err := url.PathUnescape(c.Param("id"))
		if err != nil {
			log.Fatal(err)
		}
		roomId, err := uuid.Parse(roomIdString)
		if err != nil {
			log.Fatal(err)
		}

		hub.AddClientToRoom(roomId, c)
	})
	router.GET("/xssi", func(c *gin.Context) {
		c.Data(200, "application/json", []byte(")]}'\n{\"hello\":1}"))
	})
	router.GET("/badxssi", func(c *gin.Context) {
		c.Writer.Write([]byte("sdgfsg{\"hello\":1}"))
	})
	if len(addr) > 0 {
		router.Run(addr[0])
	} else {
		router.Run()
	}
}

func serveHome(c *gin.Context) {
	f, _ := os.Open("./home.html")
	f.Close()
	fmt.Println(os.Getwd())
	fmt.Println(os.Stat("../../../../../../../cmd/gamecore/home.html"))
	c.File("../../../../../../../cmd/gamecore/home.html")
}