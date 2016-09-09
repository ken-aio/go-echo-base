package model

import (
	"time"
	"github.com/gocraft/dbr"
)

type Xxx struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewXxx(name string) *Xxx {
	return &Xxx{
		Name: name,
	}
}

func (m *Xxx) Save(tx *dbr.Tx) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	_, err := tx.InsertInto("xxxs").
		Columns("name", "created_at", "updated_at").
		Record(m).
		Exec()

	return err
}

func (m *Xxx) Load(tx *dbr.Tx, id int64) error {

	return tx.Select("*").
		From("xxxs").
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
		From("xxxs").
		Where(condition).
		LoadStruct(m)
}
