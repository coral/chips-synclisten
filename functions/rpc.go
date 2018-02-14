package functions

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/coral/chips-synclisten/chips"
	"github.com/olahol/melody"
)

type RemoteCall struct {
	Function string `json:"function"`
	Message  string `json:"message"`
}

type Response struct {
	Message string
	Data    string
}

type RPC struct {
	m      *melody.Melody
	status chan string
}

func (r *RPC) Bind(m *melody.Melody) {
	r.m = m
	r.status = make(chan string, 100)
	go r.broadcast()
	m.HandleMessage(r.HandleRemoteCall)

}

//Basically just a BROADCAST TO WEBSOCKET CHANNEL
func (r *RPC) broadcast() {
	for {
		message := <-r.status
		fmt.Println(message)
		r.m.Broadcast([]byte(message))
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

	if strings.ToLower(call.Function) == "download" {
		componumber, _ := strconv.Atoi(call.Message)
		fmt.Println("Downloading compo: ", componumber)
		r.DownloadCompo(componumber, r.status)
	}

	if strings.ToLower(call.Function) == "ping" {
		fmt.Println("PING PONG")
		r.m.Broadcast([]byte("PONG"))
	}

}

func (r *RPC) DownloadCompo(c int, status chan string) {

	compo := chips.ChipsAPI{}

	compo.LoadCompo(c)
	err := compo.DownloadCompo(status)
	if err != nil {
		fmt.Println(err)
	}
}
