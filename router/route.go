package router

import (
	"fmt"
	"github.com/MinhSang97/Server-Game/dbutil"
	middleware "github.com/MinhSang97/Server-Game/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Route routing mapping
// @title github.com/MinhSang97/Server-Game API
// @version 1.0
// @description This is a sample server github.com/MinhSang97/Server-Game server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1/api/
// @schemes http
func Route() {
	db := dbutil.ConnectDB()
	fmt.Println("Connected: ", db)

	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.SaveLogRequest())

	// Swagger endpoint
	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		api := v1.Group("/api")
		{
			api.GET("/test", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "This is a secure route"})
			})

		}
	}

	r.Run(":8080")
}
