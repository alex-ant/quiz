package questions

import (
	"encoding/json"
	"io/ioutil"
)

// Question contains single question data.
type Question struct {
	Text             string
	CorrectAnswer    string
	IncorrectAnswers []string
	ID               int
}

// Questions defines a slice of questions.
type Questions []Question

// Read reads questions from the file and returns the corresponding slice.
func Read() (Questions, error) {
	// Read file.
	data, dataErr := ioutil.ReadFile("questions.json")
	if dataErr != nil {
		return nil, dataErr
	}

	// Unmarshal the data.
	var rawQuestions map[string][]string

	jsonErr := json.Unmarshal(data, &rawQuestions)
	if jsonErr != nil {
		return nil, jsonErr
	}

	// Generate the list of questions. The first answer in the array is the
	// correct one.
	var parsedQuestions Questions
	i := 1
	for q, a := range rawQuestions {
		parsedQuestions = append(parsedQuestions, Question{
			Text:             q,
			CorrectAnswer:    a[0],
			IncorrectAnswers: a[1:],
			ID:               i, // set ID to slice index + 1 for readibility
		})

		i++
	}

	return parsedQuestions, nil
}

// GetAnswer returns the question text along with the correct answer.
func (q Questions) GetAnswer(id int) (string, string) {
	question := q[id-1]

	return question.Text, question.CorrectAnswer
}
