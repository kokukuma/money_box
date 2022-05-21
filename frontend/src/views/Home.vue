<template>
  <div class="home">
    <vs-button size="large" color="primary" type="filled" v-on:click="connect">Connect Wallet</vs-button>
    <br>
    <vs-button size="large" color="primary" type="filled" v-on:click="disconnect">Disconnect</vs-button>
    <br>
    <vs-button size="large" color="primary" type="filled" v-on:click="getBalance">GetBalance</vs-button>
    <br>
    <vs-button size="large" color="primary" type="filled" v-on:click="pray">Pray</vs-button>
    <br>
    {{accounts}}
    {{balance}}
  </div>
</template>

<script>
import Web3 from "web3";
import WalletConnectProvider from "@walletconnect/web3-provider";
import contractJSON from '../assets/MoneyBox.json'

const provider = new WalletConnectProvider({
  infuraId: "15f721c4df8c4f4f91dea73670b27d11",
  qrcodeModalOptions: {
    mobileLinks: [
      "metamask",
    ],
  },
});

const web3 = new Web3(provider);

// Subscribe to accounts change
provider.on("accountsChanged", (accounts) => {
  console.log(accounts);
});

// Subscribe to chainId change
provider.on("chainChanged", (chainId) => {
  console.log(chainId);
});

// Subscribe to session disconnection
provider.on("connect", async () => {
  const accounts = await web3.eth.getAccounts();
  alert(accounts);
});

export default {
  name: "App",
  data() {
    return {
      accounts:"",
      balance:""
    }
  },
  components: {},
  methods: {
    connect: async function () {
      await provider.enable();
      alert("connected")
      this.accounts = await web3.eth.getAccounts();
    },
    disconnect: async function () {
      await provider.disconnect();
      alert("disconnected")
    },
    getBalance: async function () {
      var myContract = new web3.eth.Contract(contractJSON, '0xa5D4B887Ac0d118D9C2C7c75918379B438E85A00', {
        from: this.accounts[0],
        // gasPrice: '20000000000' // default gas price in wei, 20 gwei in this case
      });
      const res = await myContract.methods.balance().call({
        from: this.accounts[0],
      }, function(error){
        if (error != null) {
          console.log(error)
        }
      })
      this.balance = res
    },
    pray: async function () {
      var myContract = new web3.eth.Contract(contractJSON, '0xa5D4B887Ac0d118D9C2C7c75918379B438E85A00', {
        from: this.accounts[0],
        // gasPrice: '20000000000' // default gas price in wei, 20 gwei in this case
      });

      //
      const res = await myContract.methods.pray().send({
        from: this.accounts[0],
        value: 100,
      }, function(error){
        if (error != null) {
          console.log(error)
        }
      })
      console.log(res)
      this.getBalance()
    },
  }
};
</script>
