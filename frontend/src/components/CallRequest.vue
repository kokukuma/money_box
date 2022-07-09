<template>
  <div class="call_request">
    <div>ID: <strong>{{ this.call_payload.id }}</strong></div>
    <div>Method: <strong>{{ this.call_payload.method }}</strong></div>
    <div>Params: <strong>{{ this.call_payload.params }}</strong></div>
    <br />
    <b-button @click="call_approve" variant="primary">Approve</b-button>
    <b-button @click="call_reject" variant="primary">Reject</b-button>
    <br />
  </div>
</template>

<script>
export default {
  name: 'CallRequest',
  props: [
    "call_payload",
    "connector",
    "accounts",
  ],
  data: function () {
    return {
      // call_payload: {},
    }
  },
  methods: {
    call_approve: async function() {
      const method = this.call_payload.method
      const params = this.call_payload.params
      const id = this.call_payload.id

      var Web3 = require('web3');
      var web3 = new Web3('https://ropsten.infura.io/v3/15f721c4df8c4f4f91dea73670b27d11')
      console.log(web3.eth)
      web3.eth.getGasPrice().then(console.log);

      // TODO: eth_sendTransaction
      console.log(method)
      console.log(params)
      console.log(this.accounts)
      // var gasPrice = await web3.eth.getGasPrice()
      // params[0].gas = parseInt(gasPrice * 1.3, 10)
      params[0].gas = 4712388

      this.$emit("called", id)


      const data = await this.accounts.signTransaction(params[0])
      console.log(data)

      const res = await web3.eth.sendSignedTransaction(data.rawTransaction)
      console.log(res)

      this.connector.approveRequest({
        id: id,
        result: res.transactionHash,
      });
    },
    call_reject: async function() {
      this.connector.rejectRequest({
        id: this.call_payload.id ,
        error: {
          code: "OPTIONAL_ERROR_CODE",
          message: "OPTIONAL_ERROR_MESSAGE"
        }
      });
      this.$emit("called", this.call_payload.id)
    },
  }
}
</script>

