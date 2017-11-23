package models

import (
	"time"

	"github.com/zanjs/go-echo-rest/db"
)

type (
	Product struct {
		BaseModel
		Title        string `json:"title" gorm:"type:varchar(100)"`
		ExternalCode string `json:"external_code" gorm:"varchar(100)"`
	}

	QMProduct struct {
		OwnerCode     string `json:"ownerCode"`
		ItemCode      string `json:"itemCode"`
		WarehouseCode string `json:"warehouseCode"`
		InventoryType string `json:"inventoryType"`
	}
)

func CreateProduct(m *Product) error {
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

func (m *Product) UpdateProduct(data *Product) error {
	var err error

	m.UpdatedAt = time.Now()
	m.Title = data.Title
	m.ExternalCode = data.ExternalCode

	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func (m Product) DeleteProduct() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func GetProductById(id uint64) (Product, error) {
	var (
		product Product
		err     error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Last(&product, id).Error; err != nil {
		tx.Rollback()
		return product, err
	}
	tx.Commit()

	return product, err
}

func GetProducts() ([]Product, error) {
	var (
		products []Product
		err      error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Order("id desc").Find(&products).Error; err != nil {
		tx.Rollback()
		return products, err
	}
	tx.Commit()

	return products, err
}
