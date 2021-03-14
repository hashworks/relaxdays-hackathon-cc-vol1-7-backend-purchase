package router

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gchaincl/dotsql"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Server struct {
	DotAlter   *dotsql.DotSql
	DotSelect  *dotsql.DotSql
	DB         *sql.DB
	router     *gin.Engine
	cacheStore *persistence.InMemoryStore
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// NewRouter returns a new router.
func (s Server) NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORS())

	s.cacheStore = persistence.NewInMemoryStore(time.Second)

	router.GET("/purchases", cache.CachePage(s.cacheStore, time.Hour, s.PurchaseGet))

	router.GET("/purchasesForArticle", cache.CachePage(s.cacheStore, time.Hour, s.PurchaseGetByArticleId))

	router.GET("/searchLieferant", cache.CachePage(s.cacheStore, time.Hour, s.PurchaseVendorSearch))

	router.POST("/purchase", s.PurchaseSave)

	router.GET("/purchasesBetween", cache.CachePage(s.cacheStore, time.Hour, s.PurchaseGetByTime))

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	openapiURL := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, openapiURL))

	return router
}

func (s Server) internalServerError(c *gin.Context, err string) {
	log.Print(err)
	c.Status(http.StatusInternalServerError)
}
