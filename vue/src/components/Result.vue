<template>
  <div>
      選択
        <select v-model="select_choice">
          <option value="0">Goo</option>
          <option value="1">Choki</option>
          <option value="2">Paa</option>
        </select>
        <button v-on:click="result">送信</button>
    <div>
        自分：{{ choice }}
    </div>
    <div>
        相手：グー
    </div>
    <div>
        結果：勝ち
    </div>
  </div>
</template>
<script>
import axios from 'axios'
export default {
  data () {
    return {
      choice: 'グー',
      select_choice: '0'
    }
  },
  methods: {
    result: function () {
      axios
        .get('/api/result?choice=' + this.select_choice)
        .then(response => (this.choice = response.data.PlayerSelf.choice_name))
    }
  }
}
</script>
