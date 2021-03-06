package main

import (
	"fmt"

	"github.com/coral/chips-synclisten/pkg/chips"
	"github.com/coral/chips-synclisten/pkg/credentials"
	"github.com/coral/chips-synclisten/pkg/discord"
	"github.com/coral/chips-synclisten/pkg/functions"
	"github.com/coral/chips-synclisten/pkg/polly"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

var tts polly.PollyClient
var compo chips.ChipsAPI

func main() {

	r := gin.Default()
	m := melody.New()
	cred, err := credentials.LoadCredentials()
	if err != nil {
		panic("could not load credentials")
	}
	compo = chips.ChipsAPI{}

	//WEB ROUTES

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

	//DISCORD BINDING

	discord := discord.Discord{}
	discord.SetToken(cred.Discord.Token)
	discord.SetChannel(cred.Discord.ChannelID)
	discord.Connect()

	//POLLY STUFF
	tts = polly.PollyClient{}
	tts.DefineSecrets(cred.Polly.Key, cred.Polly.Secret)
	r.POST("/tts", func(c *gin.Context) {
		message := c.PostForm("message")
		v, err := tts.GetTTS(message)
		if err != nil {
			fmt.Println(err)
			return
		}
		//vEnc := b64.StdEncoding.EncodeToString(v)
		c.Data(200, "audio/mpeg", v)
	})

	//BIND THE RPC
	rpc := functions.RPC{}
	rpc.Bind(m, &compo, &discord)

	//HOST THE WEBSERVER
	r.Run(":4020")

}
