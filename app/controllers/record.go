package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zanjs/go-echo-rest/app/models"
	"github.com/zanjs/go-echo-rest/app/services"
	"github.com/zanjs/go-echo-rest/app/utils"
)

// AllProductWareroom get all products and wareroom
func AllProductWareroom(c echo.Context) error {
	var (
		data models.ProductWareroom
		err  error
	)
	data.Products, err = models.GetProducts()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	data.Warerooms, err = models.GetWarerooms()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	go func() {

		for _, v := range data.Warerooms {
			var wareroom models.Wareroom
			wareroom = v
			numbering := wareroom.Numbering
			fmt.Println(numbering)

			for _, v2 := range data.Products {

				var product models.Product
				product = v2
				fmt.Println(product)

				var qmProduct models.QMProduct
				qmProduct.ItemCode = product.ExternalCode
				qmProduct.WarehouseCode = wareroom.Numbering
				qmProduct.OwnerCode = "bkyy"
				qmProduct.InventoryType = "ZP"

				var qmRequest models.QMRequest
				qmRequest = utils.Parameter("inventory.query", qmProduct)

				var record models.Record
				record.ProductID = product.ID
				record.WareroomID = wareroom.ID

				services.QMHTTPPost(qmRequest, record)

			}

		}

	}()

	return c.JSON(http.StatusOK, data)
}

// AllProductWareroomRecords get all products and wareroom
func AllProductWareroomRecords(c echo.Context) error {
	var (
		data     models.ProductWareroomExcel
		products []models.Product
		err      error
	)
	products, err = models.GetProducts()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	data.Warerooms, err = models.GetWarerooms()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	productExcelInt := []models.ProductExcel{}

	productExcel := productExcelInt[0:]

	for index := 2; index < len(products); index++ {
		var product models.Product
		product = products[index]

		pID := product.ID

		productExcelQuantitysInt := []models.ProductExcelQuantity{}

		productExcelQuantitys := productExcelQuantitysInt[0:]

		for _, v2 := range data.Warerooms {
			fmt.Println(v2)
			wID := v2.ID
			var record models.Record

			record, err = models.GetRecordLast(wID, pID)

			if err != nil {
				fmt.Println("查询最好一个 err: ", err, record)
			}

			var productExcelQuantity models.ProductExcelQuantity

			productExcelQuantity.Quantity = record.Quantity
			productExcelQuantity.Sales = record.Sales

			productExcelQuantitys = append(productExcelQuantitys, productExcelQuantity)
			// fmt.Println("record", record)
		}

		var pExcle models.ProductExcel
		pExcle.ProductTitle = product.Title
		pExcle.ProductExcelQuantitys = productExcelQuantitys
		productExcel = append(productExcel, pExcle)

		// data.Products = append

		// fmt.Println("Index:   Value: ", pID, index, product)
	}

	fmt.Println("productExcel:", productExcel)

	data.Products = productExcel
	// for key, v := range products {
	// 	var product models.Product
	// 	product = v
	// 	pID := product.ID
	// 	// fmt.Println(pID)
	// 	fmt.Println("key", key)

	// 	data.Products[key] = models.ProductExcel{}

	// 	for _, v2 := range data.Warerooms {
	// 		fmt.Println(v2)
	// 		wID := v2.ID
	// 		var record models.Record

	// 		record, err = models.GetRecordLast(wID, pID)

	// 		if err != nil {
	// 			fmt.Println("查询最好一个 err: ", err, record)
	// 		}

	// 		// fmt.Println("record", record)
	// 	}

	// }

	return c.JSON(http.StatusOK, data)
}

// AllRecords  get all records
func AllRecords(c echo.Context) error {
	var (
		data        models.RecordPage
		queryparams models.QueryParams
		err         error
	)

	qps := c.QueryParams()

	limitq := c.QueryParam("limit")
	offsetq := c.QueryParam("offset")
	startTimeq := c.QueryParam("start_time")
	endTime := c.QueryParam("end_time")

	limit, _ := strconv.Atoi(limitq)
	offset, _ := strconv.Atoi(offsetq)
	fmt.Println(qps)
	fmt.Println(limit)
	fmt.Println(offset)

	if limit == 0 {
		limit = 10
	}

	queryparams.Limit = limit
	queryparams.Offset = offset
	queryparams.StartTime = startTimeq
	queryparams.EndTime = endTime

	data, err = models.GetRecords(queryparams)

	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, data)
}

// ShowRecord get one record
func ShowRecord(c echo.Context) error {
	var (
		record models.Record
		err    error
	)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	record, err = models.GetRecordById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, record)
}

//DeleteRecord is record
func DeleteRecord(c echo.Context) error {
	var err error

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetRecordById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	err = m.DeleteRecord()
	return c.JSON(http.StatusNoContent, err)
}
