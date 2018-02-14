package main

import (
	"github.com/coral/chips-synclisten/rpc"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {

	r := gin.Default()
	m := melody.New()
	r.Use(static.Serve("/tmp/", static.LocalFile("tmp", true)))
	r.Use(static.Serve("/", static.LocalFile("static", true)))

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(rpc.HandleRemoteCall)

	r.Run(":4020")

}
