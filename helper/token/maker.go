package token

import (
	"time"

	"diary-app-service/entity"
)

type Maker interface {
	CreateToken(user *entity.User, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
