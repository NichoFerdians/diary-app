package http

import (
	"diary-app-service/helper/token"
	"diary-app-service/usecase"
	"log"

	"os"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(userUseCase usecase.UserUseCaseInterface, diaryUseCase usecase.DiaryUseCaseInterface) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	handlerUser := NewUserHandler(userUseCase)
	handlerDiary := NewDiaryHandler(diaryUseCase)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	SECRET := os.Getenv("JWT_SECRET")

	tokenMaker, err := token.NewJWTMaker(SECRET)
	if err != nil {
		log.Fatalf("cannot create token maker: %s", err)
	}

	v1 := router.Group("/api")
	{
		auth := v1.Group("/user")
		{
			auth.POST("/register", handlerUser.CreateUser)
			auth.POST("/login", handlerUser.LoginUser)
		}

		orders := v1.Group("/v1").Use(AuthMiddleware(tokenMaker))
		{
			orders.POST("/diary", handlerDiary.CreateDiary)
			orders.PATCH("/diary/:id", handlerDiary.UpdateDiary)
			orders.DELETE("/diary/:id", handlerDiary.DeleteDiary)
			orders.GET("/diary", handlerDiary.ListDiary)
		}
	}

	return router
}
