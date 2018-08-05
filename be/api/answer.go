package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type answerRequest []struct {
	QuestionID int    `json:"questionID"`
	Answer     string `json:"answer"`
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

	var resp []answerResponse
	for _, answer := range data {
		// Validate question ID.
		if answer.QuestionID > len(a.questions) || answer.QuestionID < 1 {
			respond("", nil, "invalid question ID received", http.StatusBadRequest, w)
			return
		}

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
