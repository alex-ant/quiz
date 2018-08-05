package api

import (
	"net/http"
)

func (a *API) getQuestionsHandler(w http.ResponseWriter, r *http.Request) {

	respond("", nil, "ok", http.StatusOK, w)
}
