package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	logging "github.com/parasit/epodcaster/pkg/log"
	models "github.com/parasit/epodcaster/pkg/models"
)

var EP_DB *gorm.DB

var Episode models.Episode
var Episodes []models.Episode

func InitDB() {
	var err error
	EP_DB, err = gorm.Open(sqlite.Open("epodcaster.db"), &gorm.Config{})
	if err != nil {
		logging.Log.Fatal(err)
	}
	EP_DB.AutoMigrate(&models.Channel{})
	EP_DB.AutoMigrate(&models.Episode{})
}

func GetChannels() []models.Channel {
	var Channels []models.Channel
	EP_DB.Model(&models.Channel{}).Preload("Episodes").Find(&Channels)
	return Channels
}

func FindChannelById(id int) models.Channel {
	var Channel models.Channel
	EP_DB.First(&Channel, id)
	return Channel
}

func AddChannel(c models.Channel) {
	result := EP_DB.Create(&c)
	logging.Log.Debugf("Id %d err %s", result.RowsAffected, result.Error)
}

func DeleteChannel(id int) {
	EP_DB.Delete(&models.Channel{}, id)
}

func GetEpisodes(channel int) []models.Episode {
	EP_DB.Where("channelID = ?", channel, &Episodes)
	return Episodes
}

func AddEpisode(episode models.Episode) {
	result := EP_DB.Create(&episode)
	logging.Log.Debugf("Id %d err %s", result.RowsAffected, result.Error)
}
