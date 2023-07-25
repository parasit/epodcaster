package main

import (
	"fmt"
	"os"
	"text/template"

	models "github.com/parasit/epodcaster/pkg/models"
	//	storage "github.com/parasit/epodcaster/pkg/storage"
	logging "github.com/parasit/epodcaster/pkg/log"
	scanner "github.com/parasit/epodcaster/pkg/tools"
)

const (
	version   = "0.0.1"
	base_link = "http://parasit.ddns.net/rss/"
)

func main() {
	fmt.Println("✓ Hello, 世界")
	//	c := channel.Channel{"Cośtam", "Ja", "Blabla", "pl-PL", base_link, []channel.Episode{
	//		{"Ep1", "Bla", time.Now(), "Bla", 1023},
	//	}}
	logging.InitLogs()
	c := models.Channel{}
	c.LoadFromFile("forta.json")
	fmt.Println("ePodcaster v." + version)
	fmt.Println(c)
	//c := storage.FindChannelById(1)
	episodes := scanner.CheckFolder(c.BaseName)
	scanner.ParseEpisodes(&c, episodes)

	t, err := template.ParseFiles("templates/channel.tpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}
