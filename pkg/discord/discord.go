package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	token   string
	channel string
	Session *discordgo.Session
}

func (d *Discord) SetToken(newToken string) {
	d.token = newToken
}

func (d *Discord) SetChannel(newChannel string) {
	d.channel = newChannel
}

func (d *Discord) Connect() {
	var err error
	d.Session, err = discordgo.New("Bot " + d.token)
	if err != nil {
		fmt.Println(err)
	}

	err = d.Session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
}

func (d *Discord) SendChannelMessage(message string) {

	d.Session.ChannelMessageSend(d.channel, message)
}
