package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/tech-rounak/true-beacon/controllers"
	middleware "github.com/tech-rounak/true-beacon/middleware"
)

func ShareRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/stock_history", controller.FetchStockHistory())
	incomingRoutes.GET("/portfolio/holdings", controller.PortfolioHoldings())
	incomingRoutes.POST("/place/order", controller.PlaceOrder())
}
