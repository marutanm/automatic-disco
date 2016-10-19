package main

import (
	"encoding/json"
	"fmt"
)

type asset struct {
	FileName string `json:"filename"`
	Md5      string `json:"md5"`
}

type assets struct {
	Contents []asset `json:"assets"`
}

func main() {
	a := new(assets)
	a.Contents = []asset{
		asset{"Images/background.jpg", "bg"},
		asset{"Images/icon.png", "icn"},
		asset{"Images/button.png", "btn"},
	}
	fmt.Println(a)
	marshaled, _ := json.Marshal(a)
	fmt.Println(string(marshaled))
}
