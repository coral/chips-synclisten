package main

import (
	"github.com/coral/chips-synclisten/pkg/chips"
	"github.com/coral/chips-synclisten/pkg/functions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {

	r := gin.Default()
	m := melody.New()
	compo := chips.ChipsAPI{}

	r.Use(static.Serve("/tmp/", static.LocalFile("tmp", true)))
	r.Use(static.Serve("/", static.LocalFile("static", true)))

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.GET("/songlist", func(c *gin.Context) {
		e := compo.GetVisualEntryList()
		c.String(200, e)
	})

	r.GET("/generatedcompo", func(c *gin.Context) {
		e := compo.GetLoadedCompo()
		c.JSON(200, e)
	})

	rpc := functions.RPC{}
	rpc.Bind(m, &compo)

	r.Run(":4020")

}
