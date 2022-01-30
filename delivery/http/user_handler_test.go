package http

import (
	"bytes"
	"diary-app-service/entity"
	"diary-app-service/usecase"
	"encoding/json"
	"os"

	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginUser(t *testing.T) {
	orig := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "testdsafknJNSKAJNDlkaDMldkmaLKNDSALDNasljdnalsdalskdmLASMDlaksdmlaskdmlak")

	mockUser := new(usecase.UserUseCaseMock)
	mockDiary := new(usecase.DiaryUseCaseMock)

	r := SetupRoutes(mockUser, mockDiary)

	loginBody := &entity.LoginUserInput{
		Username: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(loginBody)
	assert.NoError(t, err)

	mockUser.On("LoginUser", loginBody.Username, loginBody.Password).Return("jwt", nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"data\":{\"access_token\":\"jwt\"}}", w.Body.String())

	t.Cleanup(func() { os.Setenv("JWT_SECRET", orig) })
}

func TestCreateUser(t *testing.T) {
	orig := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "testdsafknJNSKAJNDlkaDMldkmaLKNDSALDNasljdnalsdalskdmLASMDlaksdmlaskdmlak")

	mockUser := new(usecase.UserUseCaseMock)
	mockDiary := new(usecase.DiaryUseCaseMock)

	r := SetupRoutes(mockUser, mockDiary)

	createBody := entity.UserInput{
		Username: "testuser",
		Password: "testpass",
		Name:     "testname",
		Email:    "testemail@mail.com",
	}

	body, err := json.Marshal(createBody)
	assert.NoError(t, err)

	user := entity.User{}

	mockUser.On("CreateUser", &createBody).Return(&user, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/user/register", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	t.Cleanup(func() { os.Setenv("JWT_SECRET", orig) })
}
