package main

import (
	"fmt"
	plexclient "github.com/akak548/go-plex-client"
)

func main() {
	fmt.Println("hello")

	plex, _ := plexclient.New("https://plex.tv", "DGaHUAmgTrTmPY_Y3UEt")
	fmt.Println(plex.GetFriends())

}
