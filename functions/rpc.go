package functions

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/coral/chips-synclisten/chips"
	"github.com/coral/chips-synclisten/messages"
	"github.com/olahol/melody"
)

type RemoteCall struct {
	Function string `json:"function"`
	Message  string `json:"message"`
}

type RPC struct {
	m      *melody.Melody
	status chan messages.RPCResponse
	compo  *chips.ChipsAPI
}

func (r *RPC) Bind(m *melody.Melody, newCompo *chips.ChipsAPI) {
	r.m = m
	r.compo = newCompo
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

	if strings.ToLower(call.Function) == "get" {
		componumber, _ := strconv.Atoi(call.Message)
		fmt.Println("Downloading compo: ", componumber)
		r.GetLoadedCompo(componumber)
	}

	if strings.ToLower(call.Function) == "download" {
		componumber, _ := strconv.Atoi(call.Message)
		fmt.Println("Downloading compo: ", componumber)
		r.DownloadCompo(componumber, r.status)
	}

	if strings.ToLower(call.Function) == "ping" {
		fmt.Println("PING PONG")
		r.m.Broadcast([]byte("PONG"))
	}

	if strings.ToLower(call.Function) == "start" {
		r.status <- messages.RPCResponse{Message: "Start", Data: call.Message}
	}

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
