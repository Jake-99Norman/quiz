package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var correctAnswers = map[string]string{
	"city":      "1",
	"year":      "0",
	"country":   "1",
	"president": "0",
	"mountain":  "2",
	"canyon":    "1",
}

func submitQuiz(c *gin.Context) {
	var submission struct {
		Answers map[string]string `json:"answers"`
	}

	if err := c.BindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	score := 0
	for question, answer := range submission.Answers {
		if correctAnswer, exists := correctAnswers[question]; exists && answer == correctAnswer {
			score++
		}
	}

	c.JSON(http.StatusOK, gin.H{"correctAnswers": score})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/submitQuiz", submitQuiz)
	router.Run("localhost:8080")
}
