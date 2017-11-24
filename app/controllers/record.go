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

// AllRecords  get all records
func AllRecords(c echo.Context) error {
	var (
		records []models.Record
		err     error
	)
	records, err = models.GetRecords()
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, records)
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
