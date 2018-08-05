package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
)

type answerRequest []struct {
	QuestionID int    `json:"questionID"`
	Answer     string `json:"answer"`
}

func (a *API) validateAnswerRequest(req answerRequest) error {
	for _, answer := range req {
		switch {
		case answer.Answer == "":
			return errors.New("empty answer")
		case answer.QuestionID > len(a.questions) || answer.QuestionID < 1:
			return errors.New("invalid questionID")
		}
	}

	return nil
}

type answerResponseQuestionsResults struct {
	QuestionText      string `json:"questionText"`
	AnsweredCorrectly bool   `json:"answeredCorrectly"`
	CorrectAnswer     string `json:"correctAnswer"`
	UserAnswer        string `json:"userAnswer"`
}

type answerResponse struct {
	CorrectAnswersPerc int                              `json:"correctAnswersPerc"`
	Percentile         int                              `json:"percentile"`
	QuestionsResults   []answerResponseQuestionsResults `json:"questionsResults"`
}

func (a *API) answerHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve request data.
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		respond("", nil, "failed to read request body", http.StatusBadRequest, w)
		return
	}

	var data answerRequest
	dataErr := json.Unmarshal(body, &data)
	if dataErr != nil {
		respond("", nil, "failed to unmarshal request body", http.StatusBadRequest, w)
		return
	}
	r.Body.Close()

	// Validate request.
	vErr := a.validateAnswerRequest(data)
	if vErr != nil {
		respond("", nil, vErr.Error(), http.StatusBadRequest, w)
		return
	}

	var res []answerResponseQuestionsResults
	var correctAnswers int
	for _, answer := range data {
		// Validate the answer.
		qText, qAnswer := a.questions.GetAnswer(answer.QuestionID)

		answeredCorrectly := answer.Answer == qAnswer
		if answeredCorrectly {
			correctAnswers++
		}

		res = append(res, answerResponseQuestionsResults{
			QuestionText:      qText,
			AnsweredCorrectly: answeredCorrectly,
			CorrectAnswer:     qAnswer,
			UserAnswer:        answer.Answer,
		})
	}

	// Calculate results.
	resultPercentage := correctAnswers * 100 / len(a.questions)
	a.results = append(a.results, resultPercentage)

	sort.Ints(a.results)

	var percentile int
	for i, r := range a.results {
		if r == resultPercentage {
			percentile = i * 100 / len(a.results)
			break
		}
	}

	// Assemble response.
	resp := answerResponse{
		QuestionsResults:   res,
		Percentile:         percentile,
		CorrectAnswersPerc: resultPercentage,
	}

	respond("results", resp, "ok", http.StatusOK, w)
}
