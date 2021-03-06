package polly

import (
	"github.com/coral/golang-tts"
)

type PollyClient struct {
	pkey    string
	psecret string
}

type PollyRequest struct {
	Message string `json:"message" binding:"required"`
}

func (p *PollyClient) GetTTS(text string) ([]byte, error) {

	polly := golang_tts.New(p.pkey, p.psecret)
	polly.Format(golang_tts.MP3)
	polly.Voice("Matthew")

	bytes, err := polly.Speech(text)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (p *PollyClient) DefineSecrets(key string, secret string) {
	p.pkey = key
	p.psecret = secret
}
