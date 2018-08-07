<template>
  <div class="container">
    <div class="row">
      <div class='hidden-xs col-sm-2 col-md-2 col-lg-2'></div>
      <div class='col-xs-12 col-sm-8 col-md-8 col-lg-8'>

        <h1>Quiz</h1>

        <br/>

        <div v-if="resultsAvailable">
          Correctly answered: {{results.correctAnswersPerc}}%
          <br/>

          You did better than {{results.percentile}}% of players
          <br/>
          <br/>
        </div>

        <div v-for="question in questions" class="container question-box" v-if="!resultsAvailable">
          <h2>{{question.text}}</h2>

          <br/>

          <div class="row">
            <div class='col-xs-6 col-sm-6 col-md-6 col-lg-6'>
              <button v-on:click="selectAnswer(question.id, question.answers[0], 0)" :id="'0' + question.id" type="button" class="btn answer-btn">{{question.answers[0]}}</button>
            </div>
            <div class='col-xs-6 col-sm-6 col-md-6 col-lg-6'>
              <button v-on:click="selectAnswer(question.id, question.answers[1], 1)" :id="'1' + question.id" type="button" class="btn answer-btn">{{question.answers[1]}}</button>
            </div>
          </div>

          <br/>

          <div class="row">
            <div class='col-xs-6 col-sm-6 col-md-6 col-lg-6'>
              <button v-on:click="selectAnswer(question.id, question.answers[2], 2)" :id="'2' + question.id" type="button" class="btn answer-btn">{{question.answers[2]}}</button>
            </div>
            <div class='col-xs-6 col-sm-6 col-md-6 col-lg-6'>
              <button v-on:click="selectAnswer(question.id, question.answers[3], 3)" :id="'3' + question.id" type="button" class="btn answer-btn">{{question.answers[3]}}</button>
            </div>
          </div>

        </div>

        <div v-for="result in results.questionsResults" class="container question-box" v-if="resultsAvailable">
          <span v-if="result.answeredCorrectly" class="correct">CORRECT</span>
          <span v-else class="incorrect">INCORRECT</span>

          <h2>{{result.questionText}}</h2>

          Your answer: {{result.userAnswer}}

          <br/>

          Correct answer: {{result.correctAnswer}}

          <br/>
        </div>

        <button v-on:click="submit()" type="button" class="btn answer-btn" v-if="!resultsAvailable">Submit answers</button>
        <button v-on:click="tryAgain()" type="button" class="btn answer-btn" v-else>Try one more time</button>

        <br/>
        <br/>

      </div>
      <div class='hidden-xs col-sm-2 col-md-2 col-lg-2'></div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

let beURL = 'http://127.0.0.1:30303';

export default {
  name: 'Questions',
  data () {
    return {
      questions: [],
      answers: {},
      results: {},
      resultsAvailable: false
    }
  },
  methods: {
    selectAnswer: function(questionID, answer, btnID) {
      this.answers[questionID] = answer;

      // Set hit button to active.
      document.getElementById(btnID.toString() + questionID).className += " active";

      // Set the rest of the buttons to inactive.
      for (let i = 0; i < 4; i++) {
        if (i !== btnID) {
          let elID = i.toString() + questionID;
          document.getElementById(elID).className = document.getElementById(elID).className.replace( /(?:^|\s)active(?!\S)/g , '' );
        }
      }
    },
    submit: function() {
      let answersReq = [];

      for (var questionID in this.answers) {
        if (this.answers.hasOwnProperty(questionID)) {
          answersReq.push({
            questionID: Number(questionID),
            answer: this.answers[questionID]
          });
        }
      }

      axios.post(beURL + '/answer', answersReq)
      .then((response) => {
        this.results = response.data.results;
        this.resultsAvailable = true;
      })
      .catch(function (error) {
        console.log(error);
      });
    },
    getQuestions: function() {
      axios.get(beURL + '/questions')
      .then(response => {
           this.questions = response.data.questions;
      })
      .catch(error => {
        console.log(error);
      })
    },
    tryAgain: function() {
      this.getQuestions();
      this.resultsAvailable = false;
      this.answers = {};
      this.results = {};
    }
  },
  created: function(){
    this.getQuestions();
  }
}
</script>

<style scoped>
.question-box {
  border-radius: 6px;
  box-shadow: 0px 2px 27px rgba(0,0,0,0.05);
  background-color: #e6f7ff;
  margin-bottom: 20px;
  padding-top: 20px;
  padding-bottom: 20px;
}
.answer-btn {
  background-color: #33bbff;
  color: #fff;
  font-family: "Montserrat";
  width: 150px;
}
.active {
  background-color: #00e600;
  color: #000;
}
.correct {
  color: #00b33c;
}
.incorrect {
  color: #cc2900;
}
</style>
