package models

var Models []interface{}

func RegisterModels() {
	Models = append(Models,
		&Account{},
		&Transaction{},
		&OperationType{},
	)
}
