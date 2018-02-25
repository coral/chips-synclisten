package chips

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/coral/chips-synclisten/pkg/messages"
	"github.com/disintegration/imaging"
	shuffle "github.com/shogo82148/go-shuffle"
)

type ChipsAPI struct {
	CompoData     CompoResponse
	FilteredCompo FilteredCompo
}

func (c *ChipsAPI) LoadCompo(compo int) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://chipscompo.com/api/compo/"+strconv.Itoa(compo), nil)
	//req, err := http.NewRequest("GET", "http://127.0.0.1:8080/46.json", nil)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Something went wrong with the request to Chips")
	} else {
		if response.StatusCode == 200 {
			resp := CompoResponse{}
			defer response.Body.Close()
			json.NewDecoder(response.Body).Decode(&resp)
			c.CompoData = resp
			c.shuffleAndSortCompo()
			return nil
		} else {
			return fmt.Errorf("404 or something from Chips")
		}

	}
	return fmt.Errorf("Could not get compo data")

}

func (c *ChipsAPI) DownloadCompo(status chan messages.RPCResponse) error {
	var p string = "tmp/compos/" + strconv.Itoa(c.CompoData.Compo.ID)
	os.MkdirAll(p, 0777)

	for _, compo := range c.CompoData.Entries {

		if compo.Type == "song" {
			status <- messages.RPCResponse{Message: "Downloading", Data: compo.Title}
			err := c.songDownloadHelper(p, compo)
			if err != nil {
				status <- messages.RPCResponse{Message: "Error", Data: err.Error()}
				fmt.Println(err)
				return err
			}
		}

	}

	for _, image := range c.CompoData.Images {

		status <- messages.RPCResponse{Message: "Downloading", Data: "image"}
		err := c.imageDownloadHelper(p, image.URL)
		if err != nil {
			status <- messages.RPCResponse{Message: "Error", Data: err.Error()}
			fmt.Println(err)
			return err
		}
	}

	status <- messages.RPCResponse{Message: "Done", Data: "Download"}

	return nil
}

func (c *ChipsAPI) GetLoadedCompo() GeneratedCompo {
	k := GeneratedCompo{
		CompoResponse: c.CompoData,
		FilteredCompo: c.FilteredCompo,
	}
	return k
}

func (c *ChipsAPI) songDownloadHelper(path string, e Entry) error {

	tokens := strings.Split(e.UploadedURL, "/")
	fileName := tokens[len(tokens)-1]

	//Check if file is already cached
	if _, err := os.Stat(path + "/" + fileName); err == nil {
		return nil
	}

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

func (c *ChipsAPI) imageDownloadHelper(path string, image string) error {
	tokens := strings.Split(image, "/")
	fileName := tokens[len(tokens)-1]

	//Create File
	output, err := os.Create(path + "/" + fileName)
	if err != nil {
		fmt.Println("Error while creating", image, fileName, "-", err)
		return fmt.Errorf("Error while creating", image, fileName, "-", err)

	}
	defer output.Close()

	//Download the compo file
	response, err := http.Get(image)
	if err != nil {
		return fmt.Errorf("Error while downloading", image, "-", err)

	}
	defer response.Body.Close()

	//Write to file
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return fmt.Errorf("Error while downloading", image, "-", err)
	}

	src, err := imaging.Open(path + "/" + fileName)
	if err != nil {
		return fmt.Errorf("Could not open image for processing", "-", err)
	}

	blurred := imaging.Blur(src, 20)
	blurred = imaging.AdjustBrightness(blurred, -10)

	err = imaging.Save(blurred, path+"/"+strings.TrimSuffix(fileName, filepath.Ext(fileName))+"_blur.jpg")
	if err != nil {
		return fmt.Errorf("Could not save processed image", "-", err)
	}

	return nil
}

func (c *ChipsAPI) GetVisualEntryList() string {
	var entrylist string
	if len(c.FilteredCompo.Songs) > 0 {
		entrylist += "---------SONGS---------" + " \n"
		for _, e := range c.FilteredCompo.Songs {
			entrylist += e.Title + " \n"
		}
		entrylist += "\n"
	}

	if len(c.FilteredCompo.Art) > 0 {
		entrylist += "---------ART---------" + " \n"
		for _, e := range c.FilteredCompo.Art {
			entrylist += e.Title + " \n"
		}
		entrylist += "\n"
	}

	if len(c.FilteredCompo.Memes) > 0 {
		entrylist += "---------MEMES---------" + " \n"
		for _, e := range c.FilteredCompo.Memes {
			entrylist += e.Title + " \n"
		}
		entrylist += "\n"
	}

	return entrylist
}

func (c *ChipsAPI) shuffleAndSortCompo() {

	//Clear out the slice
	c.FilteredCompo.Songs = nil
	c.FilteredCompo.Art = nil
	c.FilteredCompo.Memes = nil

	for _, entry := range c.CompoData.Entries {
		if entry.IsJoke {
			c.FilteredCompo.Memes = append(c.FilteredCompo.Memes, entry)
		} else {
			e := entry.Type
			switch e {
			case "song":
				c.FilteredCompo.Songs = append(c.FilteredCompo.Songs, entry)

			case "art":
				c.FilteredCompo.Art = append(c.FilteredCompo.Art, entry)

			}
		}
	}

	shuffle.Slice(c.FilteredCompo.Songs)
	shuffle.Slice(c.FilteredCompo.Art)
	shuffle.Slice(c.FilteredCompo.Memes)

}
