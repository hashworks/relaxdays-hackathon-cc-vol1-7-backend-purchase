package router

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-purchase/models"
)

func (s Server) getPurchases(purchaseRows *sql.Rows, err error, c *gin.Context) {
	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	allPurchases := make([]models.Purchase, 0)

	for purchaseRows.Next() {
		var purchase models.Purchase
		var err error
		err = purchaseRows.Scan(&purchase.Vendor, &purchase.ArticleID, &purchase.Bulk)
		if err != nil {
			s.internalServerError(c, err.Error())
			return
		}
		allPurchases = append(allPurchases, purchase)
	}

	c.JSON(http.StatusOK, allPurchases)
}

// API endpoint that returns all saved purchases
//
// @Summary Returns all saved purchases
// @Produce json
// @Success 200 {array} models.Purchase
// @Router /purchases [get]
// @Tags Purchase
func (s Server) PurchaseGet(c *gin.Context) {
	purchaseRows, err := s.DotSelect.Query(s.DB, "select-purchase")
	defer purchaseRows.Close()

	s.getPurchases(purchaseRows, err, c)
}

// API endpoint that returns all saved purchases for a given article
//
// @Summary Returns all saved purchases for a given article
// @Produce json
// @Success 200 {array} models.Purchase
// @Param x query int true "ID of article to query"
// @Router /purchasesForArticle [get]
// @Tags Purchase
func (s Server) PurchaseGetByArticleId(c *gin.Context) {
	purchaseRowsByArticleId, err := s.DotSelect.Query(s.DB, "select-purchase-by-articleId", c.Query("x"))
	defer purchaseRowsByArticleId.Close()

	s.getPurchases(purchaseRowsByArticleId, err, c)
}

// API endpoint that returns a list of vendors similar to a query
//
// @Summary Returns a list of vendors similar to a query
// @Produce json
// @Success 200 {array} models.Vendor
// @Param x query string true "Vendor query"
// @Router /searchLieferant [get]
// @Tags Purchase
func (s Server) PurchaseVendorSearch(c *gin.Context) {
	vendorRows, err := s.DotSelect.Query(s.DB, "search-vendor", c.Query("x"))
	defer vendorRows.Close()

	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	allVendors := make([]models.Vendor, 0)

	for vendorRows.Next() {
		var vendor models.Vendor
		var err error
		var _levenshteinDistance int
		err = vendorRows.Scan(&vendor, &_levenshteinDistance)
		//log.Println("levenshteinDistance of %s = %d", purchase.Vendor, _levenshteinDistance)
		if err != nil {
			s.internalServerError(c, err.Error())
			return
		}
		allVendors = append(allVendors, vendor)
	}

	c.JSON(http.StatusOK, allVendors)
}

// API endpoint that returns all saved purchases between two points in time
//
// @Summary Returns all saved purchases between two points in time
// @Produce json
// @Success 200 {array} models.Purchase
// @Failure 400 {} {} "Invalid points in time"
// @Param x query string true "Starting point in time in the format 13.03.2021 13:59:58"
// @Param y query string true "Ending point in time in the format 20.03.2021 15:59:58"
// @Router /purchasesBetween [get]
// @Tags Purchase
func (s Server) PurchaseGetByTime(c *gin.Context) {
	const layout = "02.01.2006 15:04:05"

	x_parsed, x_err := time.Parse(layout, c.Query("x"))
	y_parsed, y_err := time.Parse(layout, c.Query("y"))
	if x_err != nil || y_err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	purchaseRowsByTime, err := s.DotSelect.Query(s.DB, "search-purchase-by-time", x_parsed.Local().Unix(), y_parsed.Local().Unix())
	defer purchaseRowsByTime.Close()

	s.getPurchases(purchaseRowsByTime, err, c)
}

// API endpoint that saves a purchase
//
// @Summary Save a purchase
// @Success 204
// @Failure 400 {} {} "Invalid purchase"
// @Accept json
// @Param purchase body models.Purchase true "Purchase to save"
// @Router /purchase [post]
// @Tags Purchase
func (s Server) PurchaseSave(c *gin.Context) {
	var purchase models.Purchase
	c.BindJSON(&purchase)

	if !purchase.IsValid() {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err := s.DotAlter.Exec(s.DB, "insert-purchase", purchase.Vendor, purchase.ArticleID, purchase.Bulk)
	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	s.cacheStore.Flush()

	c.Status(http.StatusOK)
}
