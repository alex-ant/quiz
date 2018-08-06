package api

import (
	"net"
	"net/http"
	"strconv"

	"github.com/alex-ant/quiz/be/questions"
	"github.com/go-zoo/bone"
)

// API contains HTTP server's settings.
type API struct {
	port      int
	listener  net.Listener
	mux       *bone.Mux
	questions questions.Questions
	results   []int
}

// New returns new API.
func New(port int, questions questions.Questions) *API {
	return &API{
		port:      port,
		questions: questions,
	}
}

func (a *API) defineMux() {
	a.mux = bone.New()

	a.mux.Get("/questions", http.HandlerFunc(a.getQuestionsHandler))
	a.mux.Post("/answer", http.HandlerFunc(a.answerHandler))
	a.mux.Options("/answer", http.HandlerFunc(a.corsRequestHandler))
}

// Start starts the HTTP server.
func (a *API) Start() (err error) {
	a.defineMux()

	a.listener, err = net.Listen("tcp", ":"+strconv.Itoa(a.port))
	if err != nil {
		return
	}

	go http.Serve(a.listener, a.mux)

	return
}

// Stop stops the server.
func (a *API) Stop() {
	a.listener.Close()
}
