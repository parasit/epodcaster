package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/bogem/id3v2"
	"github.com/hajimehoshi/go-mp3"
	models "github.com/parasit/epodcaster/pkg/models"
	//	storage "github.com/parasit/epodcaster/pkg/storage"
)

const sampleSize = 4

func getLength(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot open %s\n", filename), err.Error())
	}
	defer f.Close()
	d, err := mp3.NewDecoder(f)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot decode %s\n", filename), err.Error())
	}
	samples := d.Length() / sampleSize             // Number of samples.
	audioLength := samples / int64(d.SampleRate()) // Audio length in seconds.
	return audioLength
}

func getTags(filename string) models.BasicTag {
	tag, err := id3v2.Open(filename, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()
	length := getLength(filename)
	newTag := models.BasicTag{FileName: filename, Artist: tag.Artist(), Title: tag.Title(), Length: length}
	return newTag
}

func CheckFolder(folder string) []models.BasicTag {
	bt := []models.BasicTag{}
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, file := range files {
		if file.Name()[len(file.Name())-3:] == "mp3" && !file.IsDir() {
			bt = append(bt, getTags(folder+"/"+file.Name()))
			log.Println(folder + "/" + file.Name())
		}
	}
	return bt
}
func defaultString(text, default_value string) string {
	if text == "" {
		return default_value
	}
	return text
}

func defaultInt64(value, default_value int64) int64 {
	if value < 1 {
		return default_value
	}
	return value
}

func ParseEpisodes(c *models.Channel, episodes []models.BasicTag) {
	for x := range episodes {
		tag := episodes[x]
		ep := models.Episode{
			ChannelID:   c.ID,
			Title:       defaultString(tag.Title, fmt.Sprintf("%s ep %d", c.BaseName, x+c.StartEpisode)),
			Description: fmt.Sprintf("Episode %d of %s", x+c.StartEpisode, c.Name),
			Length:      defaultInt64(tag.Length, 0),
			Link:        fmt.Sprintf("%s/%s", c.Link, tag.FileName),
			PubDate:     c.StartDate.Add(time.Minute * time.Duration((x+c.StartEpisode)*c.DateInterval)),
			// TODO Zrobić ładne formatowanie ścieżki, niezależnej od tego jak był podany katalog do wyszukiwania
		}
		// logging.Log.Debug(ep)
		//storage.AddEpisode(ep)
		c.Episodes = append(c.Episodes, ep)
	}
}
