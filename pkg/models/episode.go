package models

import (
	"time"

	"gorm.io/gorm"
)

type Episode struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"index:idx_ep_name_channel,unique"`
	Description string    `json:"description"`
	PubDate     time.Time `json:"pubdate"`
	Link        string    `json:"link"`
	Length      int64     `json:"length"`
	ChannelID   uint      `json:"channelid" gorm:"index:idx_ep_name_channel,unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
