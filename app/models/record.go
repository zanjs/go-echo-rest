package models

import (
	"encoding/xml"
	"time"

	"github.com/zanjs/go-echo-rest/db"
)

type (
	Record struct {
		BaseModel
		WareroomID int      `json:"wareroom_id" gorm:"type:varchar(100)"`
		ProductID  int      `json:"product_id" gorm:"type:varchar(100)"`
		Quantity   int      `json:"quantity" gorm:"type:varchar(100)"`
		Sales      int      `json:"sales" gorm:"type:varchar(100)"`
		Product    Product  `json:"product"`
		Wareroom   Wareroom `json:"wareroom"`
	}

	ProductWareroom struct {
		Products  []Product  `json:"products"`
		Warerooms []Wareroom `json:"warerooms"`
	}

	QMParameter struct {
		APPKey     string `json:"app_key"`
		CustomerID string `json:"customerid"`
		Format     string `json:"format"`
		Method     string `json:"method"`
		SignMethod string `json:"sign_method"`
		Timestamp  string `json:"timestamp"`
		Version    string `json:"v"`
	}

	QMRequest struct {
		URL  string `json:"url"`
		Body string `json:"body"`
	}

	QMResponse struct {
		XMLName xml.Name `xml:"response"`
		Flag    string   `xml:"flag"`
		Code    string   `xml:"code"`
		Message string   `xml:"message"`
		Items   []Item   `xml:"items>item"`
	}

	Items struct {
		Item Item `xml:"item"`
	}

	Item struct {
		XMLName       xml.Name `xml:"item"`
		WarehouseCode string   `xml:"warehouseCode"`
		ItemCode      string   `xml:"itemCode"`
		ItemID        string   `xml:"itemId"`
		InventoryType string   `xml:"inventoryType"`
		Quantity      int      `xml:"quantity"`
		LockQuantity  string   `xml:"lockQuantity"`
	}
)

func CreateRecord(m *Record) error {
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

func GetRecordLast(WareroomID, ProductID int) (Record, error) {
	var (
		record Record
		err    error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Where("wareroom_id = ? AND product_id = ?", WareroomID, ProductID).Find(&record).Error; err != nil {
		tx.Rollback()
		return record, err
	}
	tx.Commit()

	return record, err
}

func (m Record) DeleteRecord() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

func GetRecordById(id uint64) (Record, error) {
	var (
		record Record
		err    error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Last(&record, id).Error; err != nil {
		tx.Rollback()
		return record, err
	}
	tx.Commit()

	return record, err
}

func GetRecords() ([]Record, error) {
	var (
		records []Record
		err     error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Order("id desc").Find(&records).Error; err != nil {
		tx.Rollback()
		return records, err
	}
	tx.Commit()

	return records, err
}
