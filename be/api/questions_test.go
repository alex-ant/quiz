package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/require"
)

func TestGetQuestions(t *testing.T) {
	// Initialize test server.
	testAPI := New(0)
	testAPI.defineMux()

	testSrv := httptest.NewServer(testAPI.mux)
	defer testSrv.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  testSrv.URL,
		Reporter: httpexpect.NewRequireReporter(t),
	})

	resp := e.GET("/questions").
		Expect().
		Status(http.StatusOK).
		JSON()

	questionsArr := resp.Object().Value("questions").Array()

	var count int
	for _, v := range questionsArr.Iter() {
		obj := v.Object()

		switch obj.Value("id").Number().Raw() {
		case 1:
			obj.Value("text").Equal("q1")
			obj.Value("answers").Array().ContainsOnly("ca1", "ia1a", "ia1b", "ia1c")
		case 2:
			obj.Value("text").Equal("q2")
			obj.Value("answers").Array().ContainsOnly("ca2", "ia2a", "ia2b", "ia2c")
		case 3:
			obj.Value("text").Equal("q3")
			obj.Value("answers").Array().ContainsOnly("ca3", "ia3a", "ia3b", "ia3c")
		default:
			require.FailNow(t, "Unexpected question received")
		}

		count++
	}

	require.Equal(t, 3, count, "Invalid number of questions received")
}
