package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-purchase/models"
)

func (s Server) getPurchases(purchaseRows *sql.Rows, withLevenshteinDistance bool, err error, c *gin.Context) {
	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	allPurchases := make([]models.Purchase, 0)

	for purchaseRows.Next() {
		var purchase models.Purchase
		var err error
		if withLevenshteinDistance {
			var _levenshteinDistance int
			err = purchaseRows.Scan(&purchase.Vendor, &purchase.ArticleID, &purchase.Bulk, &_levenshteinDistance)
			//log.Println("levenshteinDistance of %s = %d", purchase.Vendor, _levenshteinDistance)
		} else {
			err = purchaseRows.Scan(&purchase.Vendor, &purchase.ArticleID, &purchase.Bulk)
		}
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

	s.getPurchases(purchaseRows, false, err, c)
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

	s.getPurchases(purchaseRowsByArticleId, false, err, c)
}

// API endpoint that returns all saved purchases with a similar vendor
//
// @Summary Returns all saved purchases with a similar vendor
// @Produce json
// @Success 200 {array} models.Purchase
// @Param x query string true "Vendor query"
// @Router /searchLieferant [get]
// @Tags Purchase
func (s Server) PurchaseGetByVendorSearch(c *gin.Context) {
	purchaseRowsByVendorSearch, err := s.DotSelect.Query(s.DB, "search-purchase-by-vendor", c.Query("x"))
	defer purchaseRowsByVendorSearch.Close()

	s.getPurchases(purchaseRowsByVendorSearch, true, err, c)
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
