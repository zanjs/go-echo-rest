package models

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/zanjs/go-echo-rest/db"
)

type (
	// RecordPage is
	RecordPage struct {
		Data []Record  `json:"data"`
		Page PageModel `json:"page"`
	}
	// Record is
	Record struct {
		BaseModel
		WareroomID int      `json:"wareroom_id" gorm:"type:varchar(100)"`
		ProductID  int      `json:"product_id" gorm:"type:varchar(100)"`
		Quantity   int      `json:"quantity" gorm:"type:varchar(100)"`
		Sales      int      `json:"sales" gorm:"type:varchar(100)"`
		Product    Product  `json:"product"`
		Wareroom   Wareroom `json:"wareroom"`
	}

	// ProductWareroom is
	ProductWareroom struct {
		Warerooms []Wareroom `json:"warerooms"`
		Products  []Product  `json:"products"`
	}

	// ProductExcel is
	ProductExcel struct {
		ProductTitle          string                 `json:"product_title"`
		ProductExcelQuantitys []ProductExcelQuantity `json:"product_excel_quantity"`
	}

	// ProductExcelQuantity is
	ProductExcelQuantity struct {
		Quantity int `json:"quantity"`
		Sales    int `json:"sales"`
	}
	// ProductWareroomExcel is
	ProductWareroomExcel struct {
		Warerooms []Wareroom     `json:"warerooms"`
		Products  []ProductExcel `json:"products"`
	}

	// QMParameter is
	QMParameter struct {
		APPKey     string `json:"app_key"`
		CustomerID string `json:"customerid"`
		Format     string `json:"format"`
		Method     string `json:"method"`
		SignMethod string `json:"sign_method"`
		Timestamp  string `json:"timestamp"`
		Version    string `json:"v"`
	}

	// QMRequest is
	QMRequest struct {
		URL  string `json:"url"`
		Body string `json:"body"`
	}

	// QMResponse is
	QMResponse struct {
		XMLName xml.Name `xml:"response"`
		Flag    string   `xml:"flag"`
		Code    string   `xml:"code"`
		Message string   `xml:"message"`
		Items   []Item   `xml:"items>item"`
	}

	//  Items is
	Items struct {
		Item Item `xml:"item"`
	}
	// Item is
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

// func GetRecords() ([]Record, error) {
// 	var (
// 		records []Record
// 		err     error
// 	)

// 	tx := gorm.MysqlConn().Begin()
// 	if err = tx.Order("id desc").Find(&records).Error; err != nil {
// 		tx.Rollback()
// 		return records, err
// 	}
// 	tx.Commit()

// 	return records, err
// }

// GetRecords is
func GetRecords(p QueryParams) (RecordPage, error) {
	var (
		records  []Record
		pageData RecordPage
		err      error
	)

	pageData.Page.Limit = p.Limit
	pageData.Page.Offset = p.Offset
	// pageData.Page.Limit = 2
	// pageData.Page.Offset = 2

	tx := gorm.MysqlConn().Begin()

	// err = tx.Find(&articles).Count(&pageData.Page.Count).Error

	// if err != nil {
	// 	return pageData, err
	// }
	timeLayout := "2006-01-02 15:04:05"

	if p.EndTime == "" {
		p.EndTime = "2099-01-01 00:00:00"
		fmt.Println("endTime 为空")
	}

	startTime, _ := time.Parse(timeLayout, p.StartTime)
	endTime, _ := time.Parse(timeLayout, p.EndTime)

	if pageData.Page.Offset == 0 {
		err = tx.Where("created_at BETWEEN ? AND ?", startTime, endTime).Preload("Wareroom").Preload("Product").Order("id desc").Limit(pageData.Page.Limit).Find(&records).Count(&pageData.Page.Count).Error
	} else {

		err = tx.Where("created_at BETWEEN ? AND ?", startTime, endTime).Preload("Wareroom").Preload("Product").Order("id desc").Offset(pageData.Page.Offset * pageData.Page.Limit).Limit(pageData.Page.Limit).Find(&records).Error
	}

	if err != nil {
		tx.Rollback()
		return pageData, err
	}

	tx.Commit()

	pageData.Data = records

	return pageData, err
}
