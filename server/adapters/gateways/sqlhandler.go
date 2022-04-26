package gateways

import (
	"gorm.io/gorm"
)

type SQLHandler interface {
	First(interface{}, ...interface{}) *gorm.DB
	Last(interface{}, ...interface{}) *gorm.DB
	Take(interface{}, ...interface{}) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
}
