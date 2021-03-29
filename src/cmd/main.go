package main

import "easy-go/src/api"

func main() {
	go api.StartApiServer(8088)
	select {}
}
