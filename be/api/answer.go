package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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

type answerResponse struct {
	QuestionText      string `json:"questionText"`
	AnsweredCorrectly bool   `json:"answeredCorrectly"`
	CorrectAnswer     string `json:"correctAnswer"`
	UserAnswer        string `json:"userAnswer"`
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

	var resp []answerResponse
	for _, answer := range data {
		// Validate the answer.
		qText, qAnswer := a.questions.GetAnswer(answer.QuestionID)

		resp = append(resp, answerResponse{
			QuestionText:      qText,
			AnsweredCorrectly: answer.Answer == qAnswer,
			CorrectAnswer:     qAnswer,
			UserAnswer:        answer.Answer,
		})
	}

	respond("results", resp, "ok", http.StatusOK, w)
}
