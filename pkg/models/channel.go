package models

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"gorm.io/gorm"
)

type Channel struct {
  ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Author       string    `json:"author"`
	Description  string    `json:"description"`
	Language     string    `json:"language"`
	Link         string    `json:"link"`
  Cover        string    `json:"cover"`
	BaseName     string    `json:"base_name"`
	StartEpisode int       `json:"start_episode"`
	StartDate    time.Time `json:"start_date"`
	DateInterval int       `json:"date_interval"` // in minutes
	Episodes     []Episode `json:"episodes"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (c *Channel) LoadFromFile(filename string) {
	file, _ := ioutil.ReadFile(filename)
	_ = json.Unmarshal([]byte(file), &c)
}
