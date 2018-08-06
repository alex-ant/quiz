<template>
  <div class="hello">
    <div class="container">
      <div class="row">
        <div class='hidden-xs col-sm-2 col-md-3 col-lg-4'></div>
        <div class='col-xs-12 col-sm-8 col-md-6 col-lg-4'>

          <div v-for="question in questions" class="container question-box">
            <h2>{{question.text}}</h2>
            
            <br/>

            <div class="row">
              <div class='col-xs-6 col-sm-6 col-md-6 col-lg-6'>
                <button v-on:click="selectAnswer(question.id, question.a1)" type="button" class="btn answer-btn">{{question.a1}}</button>
              </div>
              <div class='col-xs-6 col-sm-6 col-md-6 col-lg-6'>
                <button v-on:click="selectAnswer(question.id, question.a2)" type="button" class="btn answer-btn">{{question.a2}}</button>
              </div>
            </div>
            
            <br/>
            
            <div class="row">
              <div class='col-xs-6 col-sm-6 col-md-6 col-lg-6'>
                <button v-on:click="selectAnswer(question.id, question.a3)" type="button" class="btn answer-btn">{{question.a3}}</button>
              </div>
              <div class='col-xs-6 col-sm-6 col-md-6 col-lg-6'>
                <button v-on:click="selectAnswer(question.id, question.a4)" type="button" class="btn answer-btn">{{question.a4}}</button>
              </div>
            </div>
            
          </div>

        </div>
        <div class='hidden-xs col-sm-2 col-md-3 col-lg-4'></div>
      </div>
    </div>
  
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      questions: [
        {
          id: 1,
          text: "q1.",
          a1: "a1.",
          a2: "a2.",
          a3: "a3.",
          a4: "a4."
        },
        {
          id: 2,
          text: "q2.",
          a1: "b1.",
          a2: "b2.",
          a3: "b3.",
          a4: "b4."
        }
      ]
    }
  },
  methods: {
    selectAnswer: function(questionID, answer) {
      alert(questionID, answer);
    }
  }
}
</script>

<style scoped>
.question-box {
  border-radius: 6px;
  box-shadow: 0px 2px 27px rgba(0,0,0,0.05);
  background-color: #fff;
  margin-bottom: 20px;
  padding-top: 20px;
  padding-bottom: 20px;
}
.answer-btn {
  background-color: #3982d4;
  color: #fff;
  font-family: "Montserrat";
  width: 150px;
}
</style>
