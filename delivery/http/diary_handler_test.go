package http

import (
	"bytes"
	"diary-app-service/entity"
	"diary-app-service/helper/token"
	"diary-app-service/usecase"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDiary(t *testing.T) {
	orig := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "testdsafknJNSKAJNDlkaDMldkmaLKNDSALDNasljdnalsdalskdmLASMDlaksdmlaskdmlak")

	mockUser := new(usecase.UserUseCaseMock)
	mockDiary := new(usecase.DiaryUseCaseMock)

	r := SetupRoutes(mockUser, mockDiary)

	createBody := entity.CreateDiaryInput{
		Title:  "testTitle",
		Body:   "testBody",
		UserID: 1,
	}

	body, err := json.Marshal(createBody)
	assert.NoError(t, err)

	diary := entity.Diary{}

	mockDiary.On("CreateDiary", 1, &createBody).Return(&diary, nil)

	tokenMaker, err := token.NewJWTMaker(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Fatalf("cannot create token maker: %s", err)
	}

	user := entity.User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Name:     "Admin",
		Email:    "admin@admin.com",
	}

	accessToken, err := tokenMaker.CreateToken(&user, time.Hour)
	if err != nil {
		log.Fatalf("cannot create token : %s", err)

	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/diary", bytes.NewBuffer(body))
	reqHeader := fmt.Sprintf("Bearer %s", accessToken)

	req.Header.Set("Authorization", reqHeader)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	t.Cleanup(func() { os.Setenv("JWT_SECRET", orig) })
}

func TestUpdateDiary(t *testing.T) {
	orig := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "testdsafknJNSKAJNDlkaDMldkmaLKNDSALDNasljdnalsdalskdmLASMDlaksdmlaskdmlak")

	mockUser := new(usecase.UserUseCaseMock)
	mockDiary := new(usecase.DiaryUseCaseMock)

	r := SetupRoutes(mockUser, mockDiary)

	updateBody := entity.UpdateDiaryInput{
		Title:  "testTitleUpdate",
		Body:   "testBodyUpdate",
		UserID: 1,
	}

	body, err := json.Marshal(updateBody)
	assert.NoError(t, err)

	diary := entity.Diary{}

	mockDiary.On("UpdateDiary", 1, 1, &updateBody).Return(&diary, nil)

	tokenMaker, err := token.NewJWTMaker(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Fatalf("cannot create token maker: %s", err)
	}

	user := entity.User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Name:     "Admin",
		Email:    "admin@admin.com",
	}

	accessToken, err := tokenMaker.CreateToken(&user, time.Hour)
	if err != nil {
		log.Fatalf("cannot create token : %s", err)

	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/api/v1/diary/1", bytes.NewBuffer(body))
	reqHeader := fmt.Sprintf("Bearer %s", accessToken)

	req.Header.Set("Authorization", reqHeader)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	t.Cleanup(func() { os.Setenv("JWT_SECRET", orig) })
}

func TestDeleteDiary(t *testing.T) {
	orig := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "testdsafknJNSKAJNDlkaDMldkmaLKNDSALDNasljdnalsdalskdmLASMDlaksdmlaskdmlak")

	mockUser := new(usecase.UserUseCaseMock)
	mockDiary := new(usecase.DiaryUseCaseMock)

	r := SetupRoutes(mockUser, mockDiary)

	diary := entity.Diary{}

	mockDiary.On("DeleteDiary", 1, 1).Return(&diary, nil)

	tokenMaker, err := token.NewJWTMaker(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Fatalf("cannot create token maker: %s", err)
	}

	user := entity.User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Name:     "Admin",
		Email:    "admin@admin.com",
	}

	accessToken, err := tokenMaker.CreateToken(&user, time.Hour)
	if err != nil {
		log.Fatalf("cannot create token : %s", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/diary/1", nil)
	reqHeader := fmt.Sprintf("Bearer %s", accessToken)

	req.Header.Set("Authorization", reqHeader)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	t.Cleanup(func() { os.Setenv("JWT_SECRET", orig) })
}

func TestListDiary(t *testing.T) {
	orig := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "testdsafknJNSKAJNDlkaDMldkmaLKNDSALDNasljdnalsdalskdmLASMDlaksdmlaskdmlak")

	mockUser := new(usecase.UserUseCaseMock)
	mockDiary := new(usecase.DiaryUseCaseMock)

	r := SetupRoutes(mockUser, mockDiary)

	diary := []entity.Diary{}

	mockDiary.On("ListDiary", 1, 1, 10).Return(&diary, nil)

	tokenMaker, err := token.NewJWTMaker(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Fatalf("cannot create token maker: %s", err)
	}

	user := entity.User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Name:     "Admin",
		Email:    "admin@admin.com",
	}

	accessToken, err := tokenMaker.CreateToken(&user, time.Hour)
	if err != nil {
		log.Fatalf("cannot create token : %s", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/diary?page=1&page_size=10", nil)
	reqHeader := fmt.Sprintf("Bearer %s", accessToken)

	req.Header.Set("Authorization", reqHeader)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	t.Cleanup(func() { os.Setenv("JWT_SECRET", orig) })
}
