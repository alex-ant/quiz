package api

import (
	"github.com/alex-ant/quiz/be/questions"
)

var mockedQuestions = questions.Questions{
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
