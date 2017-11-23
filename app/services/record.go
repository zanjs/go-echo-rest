package services

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zanjs/go-echo-rest/app/models"
)

func QMHTTPPost(qmRequest models.QMRequest, record models.Record) {

	fmt.Println("qmRequest:", qmRequest)
	url := qmRequest.URL
	// url := "http://192.168.1.184:1323/"
	// url := "http://apix.jiuyescm.com/v1/qimen/receive?sign=41CAA7A7C573007A1D3C2DEBE34979B3&app_key=d8e8d76f-0917-435e-b1b3-8302282a8c4a&customerid=bkyy&format=xml&method=inventory.query&sign_method=md5&timestamp=2017-11-22 02:15:30&v=2.0"
	// var qbody = qmRequest.Body
	// var cType = "application/xml"
	// var reqBody = bytes.NewBuffer([]byte(qbody))
	// resp, err := http.NewRequest("POST", url, reqBody)
	// resp.Header.Set("Content-Type", cType)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	// handle error
	// }

	// fmt.Println("string(body)")
	// fmt.Println(string(body))

	//登陆用户名

	//json序列化
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

	item := qmResponse.Items[0].Item

	fmt.Println("xml Response Items :", item)
	fmt.Println("库存 Items :", item.Quantity)

	record.Quantity = item.Quantity

	fmt.Println("记录商品信息 Items :", record)
}
