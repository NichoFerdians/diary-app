package usecase

import (
	"diary-app-service/entity"
	"diary-app-service/helper/auth"
	"diary-app-service/helper/token"
	"diary-app-service/repository"
	"log"
	"time"
)

type UserUseCase struct {
	userRepo   repository.UserRepositoryInterface
	signingKey string
	duration   time.Duration
}

type UserUseCaseInterface interface {
	LoginUser(username string, password string) (string, error)
	CreateUser(input *entity.UserInput) (*entity.User, error)
}

func NewUserUseCase(
	userRepo repository.UserRepositoryInterface,
	signingKey string,
	duration time.Duration) *UserUseCase {
	return &UserUseCase{
		userRepo:   userRepo,
		signingKey: signingKey,
		duration:   duration,
	}
}

func (u UserUseCase) LoginUser(username string, password string) (string, error) {
	var user entity.User

	err := u.userRepo.GetUser(username, &user)
	if err != nil {
		return "", err
	}

	err = auth.CheckPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	tokenMaker, err := token.NewJWTMaker(u.signingKey)
	if err != nil {
		log.Fatalf("cannot create token maker: %s", err)
	}

	accessToken, err := tokenMaker.CreateToken(&user, u.duration)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (u UserUseCase) CreateUser(input *entity.UserInput) (*entity.User, error) {

	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	newUser := entity.User{
		Username: input.Username,
		Password: hashedPassword,
		Name:     input.Name,
		Email:    input.Email,
	}

	err = u.userRepo.CreateUser(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
