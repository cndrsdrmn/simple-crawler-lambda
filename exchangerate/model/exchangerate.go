package model

import (
	"exchangerate/dto"

	"gorm.io/gorm"
)

type ExchangeRate struct {
	gorm.Model
	BaseCurrencyCode     string           `gorm:"type:varchar(5);not null"`
	TransferCurrencyCode string           `gorm:"type:varchar(5);not null"`
	TransferDate         string           `gorm:"type:date;not null"`
	TransferRate         float64          `gorm:"not null"`
	Source               dto.ScaperSource `gorm:"type:varchar(100);not null;index"`
}
