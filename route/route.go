package route

import (
	"log"
	"shopping-cart-service/controller"
	"shopping-cart-service/middleware"
	"shopping-cart-service/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()
	httpRouter.Use(middleware.CORSMiddleware())

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	cartRepository := repository.NewCartRepo(db)
	if err := cartRepository.Migrate(); err != nil {
		log.Fatal("Cart migrate err", err)
	}

	cartController := controller.NewCartController(cartRepository)

	apiRoutes := httpRouter.Group("api/")
	{
		apiRoutes.POST("/add", cartController.AddCart)
		apiRoutes.GET("/list", cartController.ListCart)
		apiRoutes.DELETE("/delet", cartController.DeleteCart)
	}

	httpRouter.Run(":8082")
}
