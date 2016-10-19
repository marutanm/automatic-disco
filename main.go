package main

import (
	"encoding/json"
	"fmt"
)

type assetInfo struct {
	Md5 string `json:"md5"`
}

type assets struct {
	Contents map[string]assetInfo `json:"assets"`
}

func main() {
	c := map[string]assetInfo{
		"Images/background.jpg": assetInfo{"bg"},
		"Images/icon.png":       assetInfo{"icn"},
		"Images/button.png":     assetInfo{"btn"},
	}
	a := &assets{Contents: c}
	fmt.Println(a)
	marshaled, _ := json.Marshal(a)
	fmt.Println(string(marshaled))
}
