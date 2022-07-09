import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    calls: {},
  },
  mutations: {
    set_calls(state, payload) {
      state.calls[payload.id] = payload;
    },
    del_calls(state, key) {
      delete state.calls[key];
    },
  },
  actions: {
  },
  modules: {
  }
})
