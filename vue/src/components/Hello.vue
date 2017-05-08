<template>
  <section class="section">
    <div class="container ">
      <div class="notification has-text-centered">
        <h1 class="title">What time is it?</h1>
        <h2 class="subtitle">
          This app fetches the current time from a Golang backend, click the button below!
        </h2>
      </div>
      <div class="container has-text-centered">
          <a @click="getTime" class="button is-primary is-medium">Get time.Now()!</a>
      </div>
      <br>
      <div class="container  has-text-centered">
        <span id="time" class="tag is-dark is-medium">{{time}}</span>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'hello',
  data () {
    return {
      time: '--:--:--'
    }
  },
  methods: {
    getTime: function (event) {
      this.$http.get('/api').then(response => {
        this.time = response.body.time
      }, response => {
        alert('Something bad happened!')
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#time {
  width: 6em;
}
</style>
