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

type RPC struct {
	m *melody.Melody
}

func (r *RPC) Bind(m *melody.Melody) {
	r.m = m

	m.HandleMessage(r.HandleRemoteCall)
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
		DownloadCompo(componumber)
	}

	if strings.ToLower(call.Function) == "ping" {
		r.m.Broadcast([]byte("PONG"))
	}

}

func (r *RPC) downloadCompo(c int) {

	compo := chips.ChipsAPI{}

	compo.LoadCompo(c)
	err := compo.DownloadCompo()
	if err != nil {
		fmt.Println(err)
	}
}
