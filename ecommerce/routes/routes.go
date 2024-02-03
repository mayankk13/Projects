package routes

import (
	"github.com/Projects/ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproducts", controllers.AddProductsAdmin())

	incomingRoutes.GET("/users/productsview", controllers.SearchProducts())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
