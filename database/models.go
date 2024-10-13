package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email         string    `gorm:"type:text;unique;not null"`
	Password      string    `gorm:"type:text;not null"`
	UrlLinks      []UrlLink `gorm:"foreignKey:ID"`
	ReportedLinks []UrlLink `gorm:"foreignKey:ID"`
}

type UrlLink struct {
	ID          uint      `gorm:"type:int;unique;primaryKey;autoIncrement"`
	OriginalURL string    `gorm:"type:text;unique;text;not null"`
	ShortURL    string    `gorm:"type:text;not null"`
	UsedTimes   uint      `gorm:"type:int;not null"`
	CreatedAt   time.Time `gorm:"type:text;autoCreateTime;not null"`
	Reported    bool      `gorm:"type:text;default:false"`
	ReportedAt  time.Time `gorm:"type:text;default:null"`
	CreatedBy   uint      `gorm:"type:text;not null"`
	ReportedBy  []uint    `gorm:"type:int;default:null"`
}
