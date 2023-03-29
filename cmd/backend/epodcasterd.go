package main

import (
	"fmt"
	backend "github.com/parasit/epodcaster/pkg/backend"
	storage "github.com/parasit/epodcaster/pkg/storage"
  logging "github.com/parasit/epodcaster/pkg/log"
)

func main() {
	fmt.Println("Hello world")
  logging.InitLogs()
	storage.InitDB()
	backend.Run()
}
