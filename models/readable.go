package models

type readable struct {
	ID 		string `json:id`
	Type	string `json:type`
	Title 	string `json:title`
	Author	string `json:author`
	Status	string `json:status`
}