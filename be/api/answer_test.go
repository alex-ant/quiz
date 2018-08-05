package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestAnswer(t *testing.T) {
	// Initialize test server.
	testAPI := New(0, mockedQuestions)
	testAPI.defineMux()

	testSrv := httptest.NewServer(testAPI.mux)
	defer testSrv.Close()

	// Define test cases.
	cases := []struct {
		description    string
		body           []map[string]interface{}
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			description: "Correct answer",
			body: []map[string]interface{}{
				map[string]interface{}{
					"questionID": 1,
					"answer":     "ca1",
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"msg": "ok",
				"results": []map[string]interface{}{
					map[string]interface{}{
						"questionText":      "q1",
						"answeredCorrectly": true,
						"correctAnswer":     "ca1",
						"userAnswer":        "ca1",
					},
				},
				"status": http.StatusOK,
			},
		},
		{
			description: "Incorrect answer",
			body: []map[string]interface{}{
				map[string]interface{}{
					"questionID": 1,
					"answer":     "xxx",
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"msg": "ok",
				"results": []map[string]interface{}{
					map[string]interface{}{
						"questionText":      "q1",
						"answeredCorrectly": false,
						"correctAnswer":     "ca1",
						"userAnswer":        "xxx",
					},
				},
				"status": http.StatusOK,
			},
		},
		{
			description: "Mixed answers",
			body: []map[string]interface{}{
				map[string]interface{}{
					"questionID": 1,
					"answer":     "ca1",
				},
				map[string]interface{}{
					"questionID": 2,
					"answer":     "xxx",
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"msg": "ok",
				"results": []map[string]interface{}{
					map[string]interface{}{
						"questionText":      "q1",
						"answeredCorrectly": true,
						"correctAnswer":     "ca1",
						"userAnswer":        "ca1",
					},
					map[string]interface{}{
						"questionText":      "q2",
						"answeredCorrectly": false,
						"correctAnswer":     "ca2",
						"userAnswer":        "xxx",
					},
				},
				"status": http.StatusOK,
			},
		},
		{
			description: "Invalid question ID 1",
			body: []map[string]interface{}{
				map[string]interface{}{
					"questionID": -1,
					"answer":     "xxx",
				},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"msg":    "invalid questionID",
				"status": http.StatusBadRequest,
			},
		},
		{
			description: "Invalid question ID 2",
			body: []map[string]interface{}{
				map[string]interface{}{
					"questionID": 100,
					"answer":     "xxx",
				},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"msg":    "invalid questionID",
				"status": http.StatusBadRequest,
			},
		},
		{
			description: "Empty answer",
			body: []map[string]interface{}{
				map[string]interface{}{
					"questionID": 1,
					"answer":     "",
				},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"msg":    "empty answer",
				"status": http.StatusBadRequest,
			},
		},
	}

	// Run test cases.
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			e := httpexpect.WithConfig(httpexpect.Config{
				BaseURL:  testSrv.URL,
				Reporter: httpexpect.NewRequireReporter(t),
			})

			e.POST("/answer").
				WithJSON(c.body).
				Expect().
				Status(c.expectedStatus).
				JSON().
				Equal(c.expectedBody)
		})
	}
}
