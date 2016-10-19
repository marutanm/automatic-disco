package main

import (
	"encoding/json"
	"fmt"
)

type asset struct {
	FileName string `json:"filename"`
	Md5      string `json:"md5"`
}

func main() {
	assets := []asset{
		asset{"Images/background.jpg", "bg"},
		asset{"Images/icon.png", "icn"},
		asset{"Images/button.png", "btn"},
	}
	fmt.Println(assets)
	assetsB, _ := json.Marshal(assets)
	fmt.Println(string(assetsB))
}
