package models

import (
	"gorm.io/gorm"
)

// Estructura de tabla con restricciones y especificaciones y la forma en que se espera en formato json
type Task struct {
	gorm.Model

	Title       string `gorm:"type:varchar(100); not null; unique index" json:"title"`
	Description string `json:"description"`
	Done        bool   `gorm:"default: false" json:"done"`
	UserID      uint   `json:"user_id"`
}
