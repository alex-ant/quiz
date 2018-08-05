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

// Questions stores parsed list of questions and answers.
var Questions []Question

// Read reads questions from the file and returns the corresponding slice.
func Read() error {
	// Read file.
	data, dataErr := ioutil.ReadFile("questions.json")
	if dataErr != nil {
		return dataErr
	}

	// Unmarshal the data.
	var rawQuestions map[string][]string

	jsonErr := json.Unmarshal(data, &rawQuestions)
	if jsonErr != nil {
		return jsonErr
	}

	// Generate the list of questions. The first answer in the array is the
	// correct one.
	i := 1
	for q, a := range rawQuestions {
		Questions = append(Questions, Question{
			Text:             q,
			CorrectAnswer:    a[0],
			IncorrectAnswers: a[1:],
			ID:               i,
		})

		i++
	}

	return nil
}
