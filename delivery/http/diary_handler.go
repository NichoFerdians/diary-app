package http

import (
	"diary-app-service/entity"
	"diary-app-service/helper/token"
	"diary-app-service/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DiaryHandler struct {
	useCase usecase.DiaryUseCaseInterface
}

func NewDiaryHandler(useCase usecase.DiaryUseCaseInterface) *DiaryHandler {
	return &DiaryHandler{
		useCase: useCase,
	}
}

func (h *DiaryHandler) CreateDiary(c *gin.Context) {
	var input entity.CreateDiaryInput
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diary, err := h.useCase.CreateDiary(authPayload.UserID, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsp := entity.DairyResponse{
		ID:        int(diary.ID),
		Title:     diary.Title,
		Body:      diary.Body,
		CreatedAt: diary.CreatedAt,
		UpdatedAt: diary.UpdatedAt,
		DeletedAt: diary.DeletedAt,
	}

	c.JSON(http.StatusOK, gin.H{"data": rsp})
}

func (h *DiaryHandler) UpdateDiary(c *gin.Context) {
	var input entity.UpdateDiaryInput
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diary, err := h.useCase.UpdateDiary(authPayload.UserID, id, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsp := entity.DairyResponse{
		ID:        int(diary.ID),
		Title:     diary.Title,
		Body:      diary.Body,
		CreatedAt: diary.CreatedAt,
		UpdatedAt: diary.UpdatedAt,
		DeletedAt: diary.DeletedAt,
	}

	c.JSON(http.StatusOK, gin.H{"data": rsp})
}

func (h *DiaryHandler) DeleteDiary(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diary, err := h.useCase.DeleteDiary(authPayload.UserID, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsp := entity.DairyResponse{
		ID:        int(diary.ID),
		Title:     diary.Title,
		Body:      diary.Body,
		CreatedAt: diary.CreatedAt,
		UpdatedAt: diary.UpdatedAt,
		DeletedAt: diary.DeletedAt,
	}

	c.JSON(http.StatusOK, gin.H{"data": rsp})
}

func (h *DiaryHandler) ListDiary(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	diary, err := h.useCase.ListDiary(authPayload.UserID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diary})
}
