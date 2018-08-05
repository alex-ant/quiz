package api

import (
	"math/rand"
	"net/http"
	"time"
)

type getQuestionsResp struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []string `json:"answers"`
}

func (a *API) getQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	var resp []getQuestionsResp
	for _, q := range a.questions {
		// Shuffle answers.
		answers := append(q.IncorrectAnswers, q.CorrectAnswer)

		rand.Seed(time.Now().UnixNano())

		for i := range answers {
			j := rand.Intn(i + 1)
			answers[i], answers[j] = answers[j], answers[i]
		}

		// Append question to response.
		resp = append(resp, getQuestionsResp{
			ID:      q.ID,
			Text:    q.Text,
			Answers: answers,
		})
	}

	// Shuffle questions.
	rand.Seed(time.Now().UnixNano())

	for i := range resp {
		j := rand.Intn(i + 1)
		resp[i], resp[j] = resp[j], resp[i]
	}

	respond("questions", resp, "ok", http.StatusOK, w)
}
