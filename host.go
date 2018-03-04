package main

import (
	"flag"
	"fmt"

	"github.com/coral/chips-synclisten/pkg/chips"
	"github.com/coral/chips-synclisten/pkg/functions"
	"github.com/coral/chips-synclisten/pkg/polly"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

var tts polly.PollyClient
var compo chips.ChipsAPI

func main() {

	var pollyKey = flag.String("pollykey", "", "Key for AWS Polly")
	var pollySecret = flag.String("pollysecret", "", "Secret for AWS Polly")

	r := gin.Default()
	m := melody.New()
	compo = chips.ChipsAPI{}

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

	//refactor this
	tts = polly.PollyClient{}
	tts.DefineSecrets(*pollyKey, *pollySecret)
	r.POST("/tts", func(c *gin.Context) {
		message := c.PostForm("message")
		v, err := tts.GetTTS(message)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Data(200, "audio/mpeg", v)
	})

	rpc := functions.RPC{}
	rpc.Bind(m, &compo)

	r.Run(":4020")

}
