package models

import "gorm.io/gorm"

// модель для добавления в бд и для миграции
type Test_Conversion struct {
	gorm.Model
	Base_currency   string  `json:"base_currency" example:""`
	Target_currency string  `json:"target_currency" example:""`
	Amount          float64 `json:"amount" example:""`
	Result_amount   float64 `json:"result_amount" example:""`
}
