package costumers

import "github.com/google/uuid"

type Costumer struct {
	Id          uuid.UUID `json:"id" gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string    `json:"name" gorm:"column:name"`
	Document_id string    `json:"document_id" gorm:"column:document_id;unique;not null"`
	Birthdate   string    `json:"birthdate" gorm:"column:birthdate;type:date"`
}

var Costumers []Costumer
