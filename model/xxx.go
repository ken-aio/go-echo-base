package model

import (
	"github.com/gocraft/dbr"
)

type Xxx struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewXxx(name string) *Xxx {
	return &Xxx{
		Name: name,
	}
}

func (m *Xxx) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("xxx").
		Columns("name").
		Record(m).
		Exec()

	return err
}

func (m *Xxx) Load(tx *dbr.Tx, id int64) error {

	return tx.Select("*").
		From("xxx").
		Where("id = ?", id).
		LoadStruct(m)
}

type Xxxs []Xxx

func (m *Xxxs) Load(tx *dbr.Tx, name string) error {

	var condition dbr.Condition
	if name != "" {
		condition = dbr.Eq("name", name)
	}

	return tx.Select("*").
		From("xxx").
		Where(condition).
		LoadStruct(m)
}
