import Vue from 'vue'
import Router from 'vue-router'
import Result from '@/components/Result'
import Select from '@/components/Select'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Result',
      component: Result
    },
    {
      path: '/select',
      name: 'Select',
      component: Select
    }
  ]
})
