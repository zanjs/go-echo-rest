package models

import (
	"time"

	"github.com/zanjs/go-echo-rest/db"
)

type (
	Wareroom struct {
		BaseModel
		Title     string `json:"title" gorm:"type:varchar(100)"`
		Numbering string `json:"numbering" gorm:"type:varchar(100)"`
	}
)

func CreateWareroom(m *Wareroom) error {
	var err error

	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func (m *Wareroom) UpdateWareroom(data *Wareroom) error {
	var err error

	m.UpdatedAt = time.Now()
	m.Title = data.Title
	m.Numbering = data.Numbering

	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func (m Wareroom) DeleteWareroom() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func GetWareroomById(id uint64) (Wareroom, error) {
	var (
		wareroom Wareroom
		err      error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Last(&wareroom, id).Error; err != nil {
		tx.Rollback()
		return wareroom, err
	}
	tx.Commit()

	return wareroom, err
}

func GetWarerooms() ([]Wareroom, error) {
	var (
		warerooms []Wareroom
		err       error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&warerooms).Error; err != nil {
		tx.Rollback()
		return warerooms, err
	}
	tx.Commit()

	return warerooms, err
}
