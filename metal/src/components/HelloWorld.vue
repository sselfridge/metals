<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <div>
      <label for="num1"></label>
      <input type="number" name="num1" v-model="num1" />
      <label for="num1"></label>
      <input type="number" name="num1" v-model="num2" />
    </div>
    <div>
      <button id="add" class="btn selected" @click="changeOperator('add')">Add</button>
      <button id="sub" class="btn" @click="changeOperator('sub')">Subtract</button>
      <button id="mul" class="btn" @click="changeOperator('mul')">Multiply</button>
      <button id="div" class="btn" @click="changeOperator('div')">Divide</button>
    </div>
    <button @click="calculate()">Calculate!</button>
  </div>
</template>

<script>
import axios from "axios";
axios.defaults.headers.common["Access-Control-Allow-Origin"] = "*";

export default {
  name: "HelloWorld",
  data() {
    return {
      msg: "Welcome to a VUE Calculator",
      num1: 0,
      num2: 0,
      op: "add"
    };
  },
  methods: {
    changeOperator(newOp) {
      this.op = newOp;
      const btns = document.getElementsByClassName("btn");
      for (let i = 0; i < btns.length; i++) {
        const btn = btns[i];
        btn.classList.remove("selected");
        if (btn.id === newOp) {
          btn.classList.add("selected");
        }
      }
    },
    calculate() {
      const num1 = this.num1 === "" ? 0 : this.num1;
      const num2 = this.num2 === "" ? 0 : this.num2;
      const op = this.op;
      const url = `https://vfbptnnz0k.execute-api.us-west-2.amazonaws.com/meingo?num1=${num1}&num2=${num2}&op=${op}`;

      axios({ method: "GET", url: url }).then(
        result => {
          console.log(result.data);
          //TODO - data recieved as {number} r {remainder}
          this.num1 = result.data;
          this.num2 = 0;
        },
        error => {
          console.log("asdf");
          console.error(error);
        }
      );

      // axios({
      //   method: "GET",
      //   url: url,
        // headers: {
        //   "Access-Control-Allow-Origin": "*",
        //   "Access-Control-Allow-Methods": "GET",
        //   "Access-Control-Allow-Headers": "Origin, Content-Type, X-Auth-Token"
        // }
      // }).then(
      //   result => {
      //     this.ip = result.data.origin;
      //     console.log(`data: ${this.ip}`);
      //   },
      //   error => {
      //     console.error(error);
      //   }
      // );
    }
  }
};
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.selected {
  background-color: #42b983;
}
</style>
