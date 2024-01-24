package main

import (
	"log"
	"testAPI/app"
)

func main() {
	log.Println("Main log....")
	log.Fatal(app.RunAPI())
}
