package services_test

import (
	"context"
	"stackoverflow-clone/internal/services"
	"testing"

	"stackoverflow-clone/ent"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupInMemoryDB(t *testing.T) *ent.Client {
	// Initialize an SQLite in-memory database
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Fatalf("Failed to create in-memory database: %v", err)
	}

	// Run migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	return client
}

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
