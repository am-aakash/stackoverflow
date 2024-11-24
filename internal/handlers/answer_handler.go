package handlers

import (
	"net/http"
	"strconv"
	"stackoverflow-clone/internal/services"

	"github.com/gin-gonic/gin"
)

type AnswerHandler struct {
	answerService *services.AnswerService
}

func NewAnswerHandler(answerService *services.AnswerService) *AnswerHandler {
	return &AnswerHandler{answerService: answerService}
}

type CreateAnswerRequest struct {
	QuestionID int    `json:"question_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

func (h *AnswerHandler) Create(c *gin.Context) {
	userID := c.GetInt("userID")
	var req CreateAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer, err := h.answerService.CreateAnswer(c.Request.Context(), userID, req.QuestionID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, answer)
}

// Accept method to mark an answer as accepted
func (h *AnswerHandler) Accept(c *gin.Context) {
	userID := c.GetInt("userID")
	answerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid answer ID"})
			return
	}

	if err := h.answerService.AcceptAnswer(c.Request.Context(), userID, answerID); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answer accepted successfully"})
}