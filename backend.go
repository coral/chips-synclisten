package main

import (
	"fmt"
	"sync"

	"github.com/coral/chips-synclisten/chips"
)

func main() {

	var wg sync.WaitGroup

	compo := chips.ChipsAPI{}

	compo.LoadCompo(43)
	wg.Add(1)
	err := compo.DownloadCompo(&wg)
	if err != nil {
		fmt.Println(err)
	}
	wg.Wait()

}
