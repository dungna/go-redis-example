package main

import (
	"github.com/mediocregopher/radix.v2/redis"
	"log"
	"fmt"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	resp := conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	if resp.Err != nil {
		log.Fatal(resp.Err)
	}

	fmt.Println("Electrial Ladyland is added")
}
