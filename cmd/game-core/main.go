// This Source Code Form is subject to the terms of the MIT License
// See LICENSE for details
package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

var chanMap = make(openGames)

func main() {
	flag.Parse()
	hub := newHub()

	go hub.run()
	fmt.Println("hello world")
	router := gin.Default()
	router.GET("/", serveHome)
	//router.GET("/game/:id"), handleGame)
	router.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})
	router.Run()
}
func serveHome(c *gin.Context) {
	c.Data(200, "text/html", []byte(homeHTML))
}
//func handleGame(c *gin.Context) {
//
//
//	stringGameId := c.Param("id")
//	gameID, err := uuid.UUID.Parse(stringGameId)
//	if err != nil {
//		return nil, err
//	}
//	gameChan := chanMap[gameID]
//	go listenOnChan(gameChan)
//
//
//}

type openGames map[uuid.UUID]chan gameData

type gameData struct {
	user string
	time time.Time
	dataPayload string
}

//func createNewGame() {
//
//}

type Player struct {
	id uuid.UUID
	name string
}

//func (p Player) Create(name){
//	id =
//}

const homeHTML string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Chat Example</title>
    <script type="text/javascript">
        window.onload = function () {
            var conn;
            var msg = document.getElementById("msg");
            var log = document.getElementById("log");

            function appendLog(item) {
                var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
                log.appendChild(item);
                if (doScroll) {
                    log.scrollTop = log.scrollHeight - log.clientHeight;
                }
            }

            document.getElementById("form").onsubmit = function () {
                if (!conn) {
                    return false;
                }
                if (!msg.value) {
                    return false;
                }
                conn.send(msg.value);
                msg.value = "";
                return false;
            };

            if (window["WebSocket"]) {
                conn = new WebSocket("ws://" + document.location.host + "/ws");
                conn.onclose = function (evt) {
                    var item = document.createElement("div");
                    item.innerHTML = "<b>Connection closed.</b>";
                    appendLog(item);
                };
                conn.onmessage = function (evt) {
                    var messages = evt.data.split('\n');
                    for (var i = 0; i < messages.length; i++) {
                        var item = document.createElement("div");
                        item.innerText = messages[i];
                        appendLog(item);
                    }
                };
            } else {
                var item = document.createElement("div");
                item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
                appendLog(item);
            }
        };
    </script>
    <style type="text/css">
        html {
            overflow: hidden;
        }

        body {
            overflow: hidden;
            padding: 0;
            margin: 0;
            width: 100%;
            height: 100%;
            background: gray;
        }

        #log {
            background: white;
            margin: 0;
            padding: 0.5em 0.5em 0.5em 0.5em;
            position: absolute;
            top: 0.5em;
            left: 0.5em;
            right: 0.5em;
            bottom: 3em;
            overflow: auto;
        }

        #form {
            padding: 0 0.5em 0 0.5em;
            margin: 0;
            position: absolute;
            bottom: 1em;
            left: 0px;
            width: 100%;
            overflow: hidden;
        }

    </style>
</head>
<body>
<div id="log"></div>
<form id="form">
    <input type="submit" value="Send" />
    <input type="text" id="msg" size="64" autofocus />
</form>
</body>
</html>
`