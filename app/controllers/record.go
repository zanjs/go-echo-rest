package controllers

import (
	"fmt"
	"net/http"

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

	return c.JSON(http.StatusOK, data)
}
