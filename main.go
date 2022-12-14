package main

import (
	"backend-code-challenge-main/rest"
	"log"
)

func main() {
	log.Println("Main log....")
	rest.RunAPI(":3000")
}
