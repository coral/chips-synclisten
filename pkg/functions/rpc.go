package functions

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/coral/chips-synclisten/pkg/discord"

	"github.com/coral/chips-synclisten/pkg/chips"
	"github.com/coral/chips-synclisten/pkg/messages"
	"github.com/olahol/melody"
)

type RemoteCall struct {
	Function string `json:"function"`
	Message  string `json:"message"`
}

type RPC struct {
	m       *melody.Melody
	status  chan messages.RPCResponse
	compo   *chips.ChipsAPI
	discord *discord.Discord
}

func (r *RPC) Bind(m *melody.Melody, newCompo *chips.ChipsAPI, newDiscord *discord.Discord) {
	r.m = m
	r.compo = newCompo
	r.discord = newDiscord
	r.status = make(chan messages.RPCResponse, 100)
	go r.broadcast()
	m.HandleMessage(r.HandleRemoteCall)

}

//Basically just a BROADCAST TO WEBSOCKET CHANNEL
func (r *RPC) broadcast() {
	for {
		message := <-r.status
		fmt.Println(message)
		jsresp, _ := json.Marshal(message)
		r.m.Broadcast(jsresp)
	}
}

func (r *RPC) HandleRemoteCall(s *melody.Session, msg []byte) {
	call := RemoteCall{}
	err := json.Unmarshal(msg, &call)
	if err != nil {
		fmt.Println("Could not understand message: " + string(msg))
		return
	}

	//REWRITE THIS SHIT ASAP LOOOOOOOOOOOOOOOOOOOOOOOOOL

	if strings.ToLower(call.Function) == "fetch" {
		componumber, _ := strconv.Atoi(call.Message)
		fmt.Println("Fetching compo: ", componumber)
		r.FetchCompo(componumber)
	}

	if strings.ToLower(call.Function) == "get" {
		componumber, _ := strconv.Atoi(call.Message)
		fmt.Println("Returning loaded compo: ", componumber)
		r.GetLoadedCompo(componumber)
	}

	if strings.ToLower(call.Function) == "download" {
		componumber, _ := strconv.Atoi(call.Message)
		fmt.Println("Downloading compo: ", componumber)
		r.DownloadCompo(componumber, r.status)
	}

	if strings.ToLower(call.Function) == "ping" {
		fmt.Println("PING PONG")
		r.status <- messages.RPCResponse{Message: "Pong", Data: "PINGPONG"}
	}

	if strings.ToLower(call.Function) == "start" {
		r.StartCompo()
	}

	if strings.ToLower(call.Function) == "playsong" {
		songid, _ := strconv.Atoi(call.Message)
		fmt.Println(songid)
		r.PlaySong(songid)
	}

	if strings.ToLower(call.Function) == "postsonglist" {

		r.PostSongList()
	}

}

func (r *RPC) FetchCompo(c int) {
	compo := chips.ChipsAPI{}
	err := compo.LoadCompo(c)
	if err != nil {
		fmt.Println(err)
	}
	jsonCompo, _ := json.Marshal(compo.GetLoadedCompo())
	r.status <- messages.RPCResponse{Message: "Compodata", Data: string(jsonCompo)}
}

func (r *RPC) DownloadCompo(c int, status chan messages.RPCResponse) {

	r.compo.LoadCompo(c)
	err := r.compo.DownloadCompo(status)
	if err != nil {
		fmt.Println(err)
	}
}

func (r *RPC) GetLoadedCompo(c int) {
	r.compo.LoadCompo(c)
	loadedCompo := r.compo.GetLoadedCompo()
	jsonCompo, _ := json.Marshal(loadedCompo)
	r.status <- messages.RPCResponse{Message: "Compodata", Data: string(jsonCompo)}
}

func (r *RPC) StartCompo() {
	loadedCompo := r.compo.GetLoadedCompo()
	jsonCompo, _ := json.Marshal(loadedCompo)
	r.status <- messages.RPCResponse{Message: "Start", Data: string(jsonCompo)}
}

func (r *RPC) PlaySong(song int) {
	entry := r.compo.GetEntryByID(song)
	message := "NEXT UP: "
	message = message + "**" + entry.Title + "**\n<https://chipscompo.com/entry/" + strconv.Itoa(entry.ID) + ">"
	if entry.Description != "" {
		message = message + " \n```" + entry.Description + "```"
	}
	r.discord.SendChannelMessage(message)
}

func (r *RPC) PostSongList() {
	e := r.compo.GetVisualEntryList()
	r.discord.SendChannelMessage(e)
}
