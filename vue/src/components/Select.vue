<template>
  <div>
    <div>
        検索：<input type="text" v-model="word">「{{word}}」
    </div>
  <div class="list">
    <div class="item" v-for="result in results" :key="result.index">
        名前：{{result.name}}
    </div>
  </div>

  </div>
</template>
<script>
import axios from 'axios'
export default {
  data () {
    return {
      word: '',
      results: []
    }
  },
  watch: {
    word: function (newVal, oldVal) {
      console.log(newVal, oldVal)
      // this.word = newVal
      this.search(newVal)
    }
  },
  methods: {
    search: function (newVal) {
      axios
        .get('/api/result_db?word=' + newVal)
        .then(response => (this.results = response.data))
    }
  }
}
</script>

<style scoped>
.list{
  background:#A9BCD0;
    margin-top:4px;
}
.item{
  border-color: #373F51;
  border-style:solid;
  border-width:1px;
}
</style>
