package questions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var mockedQuestions = Questions{
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
}

func TestGetAnswer(t *testing.T) {
	q1, ca1 := mockedQuestions.GetAnswer(1)
	require.Equal(t, "q1", q1)
	require.Equal(t, "ca1", ca1)

	q2, ca2 := mockedQuestions.GetAnswer(1)
	require.Equal(t, "q1", q2)
	require.Equal(t, "ca1", ca2)

	q3, ca3 := mockedQuestions.GetAnswer(2)
	require.Equal(t, "q2", q3)
	require.Equal(t, "ca2", ca3)

	q4, ca4 := mockedQuestions.GetAnswer(2)
	require.Equal(t, "q2", q4)
	require.Equal(t, "ca2", ca4)
}
