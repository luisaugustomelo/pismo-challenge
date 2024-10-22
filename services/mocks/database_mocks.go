package mocks

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return &gorm.DB{Error: nil}
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(append([]interface{}{dest}, conds...)...)

	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return &gorm.DB{Error: nil}
}
