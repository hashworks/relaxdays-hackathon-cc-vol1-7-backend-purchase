package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gchaincl/dotsql"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const BasePath = "/api/v1"

type Server struct {
	DotAlter  *dotsql.DotSql
	DotSelect *dotsql.DotSql
	DB        *sql.DB
	router    *gin.Engine
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

	router.GET(BasePath+"/purchases", s.PurchaseGet)

	router.GET(BasePath+"/purchasesForArticle/:articleId", s.PurchaseGetByArticleId)

	router.POST(BasePath+"/purchase", s.PurchaseSave)

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
