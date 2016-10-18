package main

import (
	"encoding/json"
	"fmt"
)

type asset struct {
	fileName string
	md5      string
}

func main() {
	assets := []asset{
		asset{"Images/background.jpg", "..."},
		asset{"Images/icon.png", "..."},
		asset{"Images/button.png", "..."},
	}
	fmt.Println(assets)
	assetsB, _ := json.Marshal(assets)
	fmt.Println(string(assetsB))
}
