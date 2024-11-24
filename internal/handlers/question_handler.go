package handlers

import (
	"net/http"
	"stackoverflow-clone/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	questionService *services.QuestionService
	voteService     *services.VoteService
}

func NewQuestionHandler(questionService *services.QuestionService, voteService *services.VoteService) *QuestionHandler {
	return &QuestionHandler{
		questionService: questionService,
		voteService:     voteService,
	}
}

type CreateQuestionRequest struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags" binding:"required"`
}

type VoteRequest struct {
	Value int `json:"value" binding:"required,oneof=1 -1"`
}

func (h *QuestionHandler) Create(c *gin.Context) {
	userID := c.GetInt("userID") // Extract userID from middleware
	var req CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question, err := h.questionService.CreateQuestion(c.Request.Context(), userID, req.Title, req.Content, req.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, question)
}

func (h *QuestionHandler) Vote(c *gin.Context) {
	userID := c.GetInt("userID")
	questionID := c.GetInt("questionID")
	var req VoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.voteService.VoteQuestion(c.Request.Context(), userID, questionID, req.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *QuestionHandler) Search(c *gin.Context) {
	query := c.Query("q")
	tags := c.QueryArray("tags")

	questions, err := h.questionService.SearchQuestions(c.Request.Context(), query, tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (h *QuestionHandler) GetQuestionWithAnswers(c *gin.Context) {
	questionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	questionWithAnswers, err := h.questionService.GetQuestionWithAnswers(c.Request.Context(), questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questionWithAnswers)
}
