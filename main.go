package main

import (
	"diary-app-service/delivery/http"
	"diary-app-service/helper/db"
	"diary-app-service/repository"
	"diary-app-service/usecase"
	"os"
	"time"
)

func main() {
	db := db.Connect()

	userRepo := repository.NewUserRepository(db)
	diaryRepo := repository.NewDiaryRepository(db)

	userUseCase := usecase.NewUserUseCase(
		userRepo,
		os.Getenv("JWT_SECRET"),
		time.Hour,
	)

	diaryUseCase := usecase.NewDiaryUseCase(
		diaryRepo,
	)

	r := http.SetupRoutes(*userUseCase, *diaryUseCase)
	r.Run()
}
