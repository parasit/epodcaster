package main

import (
	"os"
	"text/template"

	scanner "github.com/parasit/epodcaster/pkg/tools"
	storage "github.com/parasit/epodcaster/pkg/storage"
)

const (
	version   = "0.0.1"
	base_link = "http://parasit.ddns.net/rss/"
)

func main() {
	//	c := channel.Channel{"Cośtam", "Ja", "Blabla", "pl-PL", base_link, []channel.Episode{
	//		{"Ep1", "Bla", time.Now(), "Bla", 1023},
	//	}}
	// c := models.Channel{}
	// c.LoadFromFile("test.json")
	// fmt.Println("ePodcaster v." + version)
	// fmt.Println(c)
	c := storage.FindChannelById(1)
	episodes := scanner.CheckFolder("vm1")
	scanner.ParseEpisodes(c, episodes)

	t, err := template.ParseFiles("templates/channel.tpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}
