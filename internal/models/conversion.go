package models

import "gorm.io/gorm"

type Test_Conversion struct {
	gorm.Model
	Base_currency   string  `json:"base_currency" example:""`
	Target_currency string  `json:"target_currency" example:""`
	Amount          float64 `json:"amount" example:""`
}
