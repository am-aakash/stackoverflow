package services_test

import (
	"context"
	"stackoverflow-clone/internal/services"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuestion(t *testing.T) {
	client := setupInMemoryDB(t)
	defer client.Close()

	// Setup: Create a user for the test
	user := client.User.Create().
		SetUsername("testuser").
		SetEmail("testuser@example.com").
		SetPasswordHash("hashedpassword").
		SetDisplayName("Test User").
		SaveX(context.Background())

	// Create QuestionService
	questionService := services.NewQuestionService(client)

	// Action: Create a question with tags
	tags := []string{"Go", "Programming"}
	question, err := questionService.CreateQuestion(context.Background(), user.ID, "Test Title", "Test Content", tags)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, question)
	assert.Equal(t, "Test Title", question.Title)
	assert.Equal(t, user.ID, question.Edges.Author.ID)

	// Verify: Check if tags were created and associated
	assert.Len(t, question.Edges.Tags, 2)
	assert.Contains(t, question.Edges.Tags[0].Name, "Go")
}
