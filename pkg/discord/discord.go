package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	token   string
	session *discordgo.Session
}

func (d *Discord) SetToken(newToken string) {
	d.token = newToken
}

func (d *Discord) Connect() {
	var err error
	d.session, err = discordgo.New("Bot " + d.token)
	if err != nil {
		fmt.Println(err)
	}

	err = d.session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
}
