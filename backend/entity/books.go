package entity

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title	string		`valid:"required~Title is required;stringlength(3|100)~Title must be between 3 and 100"`
	Price	float64		`valid:"required~Price is required;float~Price must be float;range(50.00|5000.00)~Prince must be between 50 and 5000"`
	Code	string		`valid:"required~Code is required;matches(^[B][K]\\d{6}$)~Code must be start with BK followed by 6 digits (0-9)"`
}