package main

import (
	"fmt"
	"github.com/skybase22/test-go-api/router"
)

func main() {
	stringHello := "Hello Jook."
	fmt.Println(stringHello)
	router.New(8000).Start()
}
