<template>
  <div class="wallet">
    <div v-if="!accountExists">
      <b-button @click="createAccount" variant="danger">Create Account</b-button>
    </div>

    <div v-if="accountExists">
      <div>Account: <strong>{{ accounts.address }}</strong></div>
      <br />
      <b-button variant="primary">Send Transaction</b-button>
      <b-button variant="primary" v-b-modal.modal-scan>Scan</b-button>
      <br />
      <br />
 
      <!-- Table -->
      <div v-if="connectedSites.length!==0">
        <b-table striped hover :items="connectedSites">
          <template v-slot:cell(disconnect)="data">
            <b-button @click="kill(data.item.name)" variant="primary" v-if="data.item.disconnect" >disconnect</b-button>
          </template>
        </b-table>
      </div>

      <!-- Table -->
      <div v-if="call_candidate.length!==0">
        <b-table striped hover :items="call_candidate">
          <template v-slot:cell(payload)="data">
            <b-button @click="call(data.item.payload)" variant="primary">call</b-button>
          </template>
        </b-table>
      </div>
    </div>

    <b-modal id="modal-scan" title="Scan">
      <div v-if="step === 'form'">
        <b-form>
          <b-form-group label="URL:" label-for="input-2">
            <b-form-input
              id="input-2"
              v-model="wc_url"
              placeholder="wc:******"
              required
            ></b-form-input>
          </b-form-group>
          <br />
          <b-button @click="submit" variant="primary">Submit</b-button>
        </b-form>
      </div>

      <div v-if="step === 'approval'">
        <div>Name: <strong>{{ candidateSites.name }}</strong></div>
        <div>URL: <strong>{{ candidateSites.url }}</strong></div>
        <div>Description: <strong>{{ candidateSites.description }}</strong></div>
        <br />
        <b-button @click="approve" variant="primary">Approve</b-button>
        <b-button @click="reject" variant="primary">Reject</b-button>
        <br />
      </div>
    </b-modal>

    <b-modal id="modal-call" title="Scan">
      <CallRequest
        :call_payload="call_payload"
        :connector="connector"
        :accounts="accounts"
        @called="handleCalled"
      />
      <!-- <div>ID: <strong>{{ this.call_payload.id }}</strong></div> -->
      <!-- <div>Method: <strong>{{ this.call_payload.method }}</strong></div> -->
      <!-- <div>Params: <strong>{{ this.call_payload.params }}</strong></div> -->
      <!-- <br /> -->
      <!-- <b&#45;button @click="call_approve" variant="primary">Approve</b&#45;button> -->
      <!-- <b&#45;button @click="call_reject" variant="primary">Reject</b&#45;button> -->
      <!-- <br /> -->
    </b-modal>

  </div>
</template>

<script>
import WalletConnect from "@walletconnect/client";
import CallRequest from '@/components/CallRequest.vue'

export default {
  name: "Wallet",
  data() {
    return {
      accountExists: false,
      wc_url: "",
      accounts: "",
      connector: "",
      step: "form",
      connectedSites: [],
      candidateSites: "",
      call_payload: {},
      call_candidate: [],
    }
  },
  components: {
    CallRequest,
  },
  created: function () {
    this.accounts = this.getAccount()
    if (this.accounts != null) {
      this.accountExists = true
    }
    console.log(this.accounts)

    const local = localStorage.getItem("walletconnect")
    if (local != null) {
      const session = JSON.parse(local)
      this.connector = new WalletConnect({session});
      this.setEvent(this.connector)
      this.updateConnectedSites(session.peerMeta)
    }
    console.log(this.connector)
  },
  methods: {
    getAccount: function() {
      var Web3 = require('web3');
      var web3 = new Web3('https://ropsten.infura.io/v3/15f721c4df8c4f4f91dea73670b27d11')
      const privateKey = localStorage.getItem("privatekey")
      let accounts = web3.eth.accounts.privateKeyToAccount(privateKey);
      return accounts
    },
    createAccount: async function() {
      var Accounts = require('web3-eth-accounts');
      var accounts = new Accounts('https://ropsten.infura.io/v3/15f721c4df8c4f4f91dea73670b27d11');
      var account =  accounts.create()
      localStorage.setItem("address", account.address)
      localStorage.setItem("privatekey", account.privateKey)
      console.log(account)
      return account
    },
    approve: async function() {
      this.accounts = this.getAccount()
      console.log("--------------- ap")
      console.log(this.accounts)

      this.connector.approveSession({
        accounts: [
          this.accounts.address,
        ],
        chainId: 3
      })
      this.step = "approved"
      this.$bvModal.hide("modal-scan")
      this.updateConnectedSites(this.candidateSites)

    },
    updateConnectedSites(peerMeta) {
      this.connectedSites.push({
        icon: peerMeta.icons,
        name: peerMeta.name,
        url: peerMeta.url,
        description: peerMeta.description,
        disconnect: true,
      })
    },
    reject: async function() {
      this.connector.rejectSession({
        message: 'OPTIONAL_ERROR_MESSAGE'
      })
      this.step = "form"
    },
    kill: async function() {
      this.connector.killSession()
      this.connectedSites = []
    },
    showApproval: async function(peerMeta) {
      this.step = "approval"
      this.candidateSites = peerMeta
    },
    call: async function(payload) {
      console.log(payload)
      this.call_payload = payload

      this.$bvModal.show("modal-call")
    },
    handleCalled: async function(id) {
      this.$bvModal.hide("modal-call")
      this.$store.commit('del_calls', id)
      this.updateCallCandidate()
    },
    submit: async function () {
      console.log(this.wc_url)
      this.connector = new WalletConnect({
        uri: this.wc_url,
        clientMeta: {
          description: "Wallet",
          name: "Wallet",
        },
      });
      this.setEvent(this.connector)
      console.log(this.connector)
    },
    updateCallCandidate: async function() {
      this.call_candidate = []
      for (let i=0; i < Object.keys(this.$store.state.calls).length; i++){
        let key = Object.keys(this.$store.state.calls)[i]
        let payload = this.$store.state.calls[key]
        this.call_candidate.push({
          id: payload.id,
          method: payload.method,
          payload: payload,
        })
      }
    },
    setEvent(connector) {
      connector.on("session_request", (error, payload) => {
        if (error) {
          throw error;
        }
        console.log(payload)
        this.showApproval(payload.params[0].peerMeta)
      });
      connector.on("call_request", (error, payload) => {
        if (error) {
          throw error;
        }
        console.log("------------- call_request")
        this.$store.commit('set_calls', payload)
        console.log(this.$store.state.calls)
        this.updateCallCandidate()
      });
      connector.on("disconnect", (error, payload) => {
        if (error) {
          throw error;
        }
        console.log(payload)
        this.step = "form"
      });
    },
  }
};
</script>
