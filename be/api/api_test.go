package api

import (
	"os"
	"testing"

	"github.com/alex-ant/quiz/be/questions"
)

func TestMain(m *testing.M) {
	// Mock questions.
	questions.Questions = []questions.Question{
		{
			Text:          "q1",
			CorrectAnswer: "ca1",
			IncorrectAnswers: []string{
				"ia1a",
				"ia1b",
				"ia1c",
			},
			ID: 1,
		},
		{
			Text:          "q2",
			CorrectAnswer: "ca2",
			IncorrectAnswers: []string{
				"ia2a",
				"ia2b",
				"ia2c",
			},
			ID: 2,
		},
		{
			Text:          "q3",
			CorrectAnswer: "ca3",
			IncorrectAnswers: []string{
				"ia3a",
				"ia3b",
				"ia3c",
			},
			ID: 3,
		},
	}

	os.Exit(m.Run())
}
