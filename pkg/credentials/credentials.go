package credentials

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Credentials struct {
	Polly struct {
		Key    string `json:"key"`
		Secret string `json:"secret"`
	} `json:"polly"`
}

func LoadCredentials() (Credentials, error) {
	jsonFile, err := os.Open("credentials.json")
	if err != nil {
		return Credentials{}, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cred Credentials
	json.Unmarshal(byteValue, &cred)

	return cred, nil
}
