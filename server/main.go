package main

import (
	"context"
	"log"
	"stackoverflow-clone/ent"
	"stackoverflow-clone/internal/config"
	"stackoverflow-clone/internal/handlers"
	"stackoverflow-clone/internal/middleware"
	"stackoverflow-clone/internal/services"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize database connection
	db, err := sql.Open("mysql", cfg.GetDBConnString())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create ent client
	client := ent.NewClient(ent.Driver(db))
	defer client.Close()

	// Run schema migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	// Initialize services
	userService := services.NewUserService(client)
	questionService := services.NewQuestionService(client)
	voteService := services.NewVoteService(client)
	answerService := services.NewAnswerService(client)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService, cfg.JWTSecret)
	questionHandler := handlers.NewQuestionHandler(questionService, voteService)
	answerHandler := handlers.NewAnswerHandler(answerService)

	// Initialize Gin router
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Public routes
	r.POST("/api/register", userHandler.Register)
	r.POST("/api/login", userHandler.Login)
	r.GET("/api/questions", questionHandler.Search)
	r.GET("/api/questions/:id", questionHandler.GetQuestionWithAnswers)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		protected.POST("/questions", questionHandler.Create)
		protected.POST("/questions/:id/vote", questionHandler.Vote)
		protected.POST("/answers", answerHandler.Create)
		protected.PATCH("/answers/:id/accept", answerHandler.Accept)
	}

	// Start server
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
