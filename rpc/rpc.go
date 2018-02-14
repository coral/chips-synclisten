package rpc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coral/chips-synclisten/chips"
	"github.com/olahol/melody"
)

type RemoteCall struct {
	Function string `json:"function"`
	Message  string `json:"message"`
}

func HandleRemoteCall(s *melody.Session, msg []byte) {
	call := RemoteCall{}
	err := json.Unmarshal(msg, &call)
	if err != nil {
		fmt.Println("Could not understand message: " + string(msg))
		return
	}

	if strings.ToLower(call.Function) == "download" {

		DownloadCompo(43)
	}

}

func DownloadCompo(c int) {

	compo := chips.ChipsAPI{}

	compo.LoadCompo(c)
	err := compo.DownloadCompo()
	if err != nil {
		fmt.Println(err)
	}
}
