package services

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zanjs/go-echo-rest/app/models"
	"github.com/zanjs/go-echo-rest/app/utils"
)

func QMHTTPPost(qmRequest models.QMRequest, record models.Record) {

	fmt.Println("qmRequest:", qmRequest)
	url := qmRequest.URL

	post := qmRequest.Body

	fmt.Println("post:? ", post)

	fmt.Println("URL:>", url)

	var xmlStr = []byte(post)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	// boydStr := string(body)

	qmResponse := models.QMResponse{}

	// body2 := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	// <response>
	// 	<flag>success</flag>
	// 	<code>SUCCESS</code>
	// 	<message>查询成功!</message>
	// 	<items>
	// 		<warehouseCode>B01</warehouseCode>
	// 		<itemCode>1371937585362246455</itemCode>
	// 		<itemId>1700046145</itemId>
	// 		<inventoryType>CC</inventoryType>
	// 		<quantity>3</quantity>
	// 		<lockQuantity>0</lockQuantity>
	// 	</items>
	// </response>`

	err = xml.Unmarshal([]byte(body), &qmResponse)
	if err != nil {
		fmt.Println("xml sp err :", err)
	}

	code := qmResponse.Code

	if code != "SUCCESS" {
		return
	}

	item := qmResponse.Items[0]

	fmt.Println("xml Response Items :", item)
	fmt.Println("库存 Items :", item.Quantity)

	fmt.Println("记录商品信息 Items :", record)

	oldRecord, oErr := models.GetRecordLast(record.WareroomID, record.ProductID)

	if oErr != nil {
		fmt.Println("查询旧数据：err:", oErr)
	}

	fmt.Println("查询旧数据：", oldRecord)

	var initSales = 0
	var quantity = item.Quantity
	var oldQuantity = oldRecord.Quantity

	if utils.IsEmpty(oldQuantity) {
		fmt.Println("旧库存：", oldQuantity)
	}

	fmt.Println("旧库存：", oldQuantity)
	fmt.Println("新库存：", quantity)

	if oldQuantity == quantity {
		return
	}

	if oldQuantity > quantity {
		initSales = oldQuantity - quantity
	}

	// if quantity >= oldQuantity {

	// }

	newRecord := new(models.Record)

	newRecord.Quantity = item.Quantity
	newRecord.Sales = initSales
	newRecord.ProductID = record.ProductID
	newRecord.WareroomID = record.WareroomID

	models.CreateRecord(newRecord)
}
