# backend

### <a name="execution_flags"></a>Execution Flags

|Flag|Env. variable|Default value|Description|
|:----|:----|:---|:---|
|api-port|API_PORT|30303|HTTP API port number|

### Running tests

`go test -cover ./...`

### Adding questions

The questions reside in `questions.json` file. Each question is represented as a map key along with an array of possible answers in the form of the corresponding value. The first answer in the array must be the correct one.

### API

* [GET /questions](#get_questions) - Get a shuffled list of questions.
* [POST /answer](#post_answer) - Submit answers and get the results.

#### <a name="get_questions"></a>GET /questions

*Get a shuffled list of questions.*

Sample request:

```
curl 'http://localhost:30303/questions'
```

Success response:

```json
{
  "msg": "ok",
  "questions": [
    {
      "id": 1,
      "text": "Where would you find the Sea of Tranquility?",
      "answers": [
        "The USA",
        "The Moon",
        "Canada",
        "Australia"
      ]
    },
    {
      "id": 3,
      "text": "Name the world's biggest island.",
      "answers": [
        "Iceland",
        "Greenland",
        "Australia",
        "Sicily"
      ]
    },
    {
      "id": 2,
      "text": "Name the seventh planet from the sun.",
      "answers": [
        "Saturn",
        "Nibiru",
        "Jupiter",
        "Uranus"
      ]
    }
  ],
  "status": 200
}
```

Both the list of questions and each question's answers are shuffled upon each request.

#### <a name="post_answer"></a>POST /answer

*Submit answers and get the results.*

Sample request:

```
curl 'http://localhost:30303/answer' -d '[
  {
    "questionID": 1,
    "answer": "Amazon"
  },
  {
    "questionID": 2,
    "answer": "Australia"
  }
]'
```

Success response:

```json
{
  "msg": "ok",
  "results": {
    "percentile": 33,
    "correctAnswersPerc": 50,
    "questionsResults": [
      {
        "questionText": "What is the world's longest river?",
        "answeredCorrectly": true,
        "correctAnswer": "Amazon",
        "userAnswer": "Amazon"
      },
      {
        "questionText": "Name the world's biggest island.",
        "answeredCorrectly": false,
        "correctAnswer": "Greenland",
        "userAnswer": "Australia"
      }
    ]
  },
  "status": 200
}
```