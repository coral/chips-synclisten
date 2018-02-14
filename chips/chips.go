package chips

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ChipsAPI struct {
	CompoData CompoResponse
}

func (c *ChipsAPI) LoadCompo(compo int) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://chipscompo.com/api/compo/"+strconv.Itoa(compo), nil)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Something went wrong with the request to Chips")
	} else {
		if response.StatusCode == 200 {
			resp := CompoResponse{}
			defer response.Body.Close()
			json.NewDecoder(response.Body).Decode(&resp)
			c.CompoData = resp
			return nil
		} else {
			return fmt.Errorf("404 or something from Chips")
		}

	}
	return fmt.Errorf("Could not get compo data")

}

func (c *ChipsAPI) DownloadCompo() error {
	var p string = "tmp/compos/" + strconv.Itoa(c.CompoData.Compo.ID)
	os.MkdirAll(p, 0777)

	for _, compo := range c.CompoData.Entries {

		if compo.Type == "song" {
			err := c.downloadHelper(p, compo)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}

	}

	return nil
}

func (c *ChipsAPI) downloadHelper(path string, e Entry) error {

	fmt.Println("Downloading: " + e.Title)

	tokens := strings.Split(e.UploadedURL, "/")
	fileName := tokens[len(tokens)-1]

	//Create File
	output, err := os.Create(path + "/" + fileName)
	if err != nil {
		fmt.Println("Error while creating", e.Title, e.UploadedURL, fileName, "-", err)
		return fmt.Errorf("Error while creating", e.Title, e.UploadedURL, fileName, "-", err)

	}
	defer output.Close()

	//Download the compo file
	response, err := http.Get(e.UploadedURL)
	if err != nil {
		return fmt.Errorf("Error while downloading", e.UploadedURL, "-", err)

	}
	defer response.Body.Close()

	//Write to file
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return fmt.Errorf("Error while downloading", e.UploadedURL, "-", err)
	}

	return nil
}
