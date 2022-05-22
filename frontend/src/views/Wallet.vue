<template>
  <div class="wallet">
    <div v-if="step === 'form'">
      <div class="centerx labelx">
        <vs-input label="URL" placeholder="wc:*****" v-model="wc_url"/>
      </div>
      <vs-button size="large" color="primary" type="filled" v-on:click="submit">Submit</vs-button>
    </div>
    <div v-if="step === 'approval'">
      {{app_name}}
      <vs-button size="large" color="primary" type="filled" v-on:click="approve">Approve</vs-button>
      <vs-button size="large" color="primary" type="filled" v-on:click="reject">Reject</vs-button>
    </div>
    <div v-if="step === 'approved'">
      Connected
      {{app_name}}
      <vs-button size="large" color="primary" type="filled" v-on:click="kill">Disconnect</vs-button>
    </div>
  </div>
</template>

<script>
import WalletConnect from "@walletconnect/client";

export default {
  name: "Wallet",
  data() {
    return {
      wc_url: "",
      accounts: "",
      connector: "",
      step: "form",
      app_name: "",
    }
  },
  components: {},
  methods: {
    approve: async function() {
      this.connector.approveSession({
        accounts: [
          '0xee6983ACaCe98e766E5b78f0533EA74c6e734dE5',
        ],
        chainId: 3
      })
      this.step = "approved"
    },
    reject: async function() {
      this.connector.rejectSession({
        message: 'OPTIONAL_ERROR_MESSAGE'
      })
      this.step = "form"
    },
    kill: async function() {
      this.connector.killSession()
    },
    showApproval: async function(peerMeta) {
      // TODO: 鍵作成
      this.step = "approval"
      this.app_name = peerMeta
    },
    call: async function(method, params) {
      // TODO: eth_sendTransaction
      console.log(method)
      console.log(params)

      // // Approve Call Request
      // this.connector.approveRequest({
      //   id: 1,
      //   result: "0x41791102999c339c844880b23950704cc43aa840f3739e365323cda4dfa89e7a"
      // });
      //
      // // Reject Call Request
      // this.connector.rejectRequest({
      //   id: 1,                                  // required
      //   error: {
      //     code: "OPTIONAL_ERROR_CODE"           // optional
      //     message: "OPTIONAL_ERROR_MESSAGE"     // optional
      //   }
      // });
    },
    submit: async function () {
      console.log(this.wc_url)
      this.connector = new WalletConnect(
        {
          // Required
          uri: this.wc_url,
          // Required
          clientMeta: {
            description: "WalletConnect Developer App",
            url: "https://walletconnect.org",
            icons: ["https://walletconnect.org/walletconnect-logo.png"],
            name: "WalletConnect",
          },
        }
      );

      this.connector.on("session_request", (error, payload) => {
        if (error) {
          throw error;
        }
        console.log(payload)
        this.showApproval(payload.params[0].peerMeta)
      });
      this.connector.on("call_request", (error, payload) => {
        if (error) {
          throw error;
        }
        this.call(payload.method, payload.params)
      });
      this.connector.on("disconnect", (error, payload) => {
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
