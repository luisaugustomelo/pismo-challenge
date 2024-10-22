package interfaces

import "gorm.io/gorm"

type Database interface {
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	//Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	//Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	//Save(value interface{}) (tx *gorm.DB)
}
