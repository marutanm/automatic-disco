package main

import "fmt"

type asset struct {
	fileName string
	md5      string
}

func main() {
	asset := new(asset)
	asset.fileName = "this_is_file_name"
	asset.md5 = "abcdefg"
	fmt.Println(asset)
}
