package models

import "encoding/xml"

type (
	Record struct {
		BaseModel
		WareroomID int `json:"wareroom_id" gorm:"type:varchar(100)"`
		ProductID  int `json:"product_id" gorm:"type:varchar(100)"`
		Quantity   int `json:"quantity" gorm:"type:varchar(100)"`
		Sales      int `json:"sales" gorm:"type:varchar(100)"`
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
		Items   []Items  `xml:"items"`
	}

	Items struct {
		Item Item `xml:"item"`
	}

	Item struct {
		// XMLName       xml.Name `xml:"item"`
		WarehouseCode string `xml:"warehouseCode"`
		ItemCode      string `xml:"itemCode"`
		ItemID        string `xml:"itemId"`
		InventoryType string `xml:"inventoryType"`
		Quantity      int    `xml:"quantity"`
		LockQuantity  string `xml:"lockQuantity"`
	}
)
