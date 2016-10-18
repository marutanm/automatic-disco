package main

import "fmt"

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
}
