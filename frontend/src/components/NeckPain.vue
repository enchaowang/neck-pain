<script setup>
import { reactive } from 'vue'
import { Start,Stop } from '../../wailsjs/go/main/App'

// images
import img_neck_pain_1 from "../assets/images/neck-pain-1.png"
import img_neck_pain_ok1 from "../assets/images/neck-pain-ok1.png"
import img_neck_pain_ok2 from "../assets/images/neck-pain-ok2.png"

// images for show
const img_neck_pains = [img_neck_pain_1]
const img_neck_pains_ok = [img_neck_pain_ok1, img_neck_pain_ok2]

const data = reactive({
  status: "on",
  btnStr: "Start!",
  mode: "random",
  logoImage: img_neck_pain_1,
})

function btnfunc() {
  if (data.status == "on") {
    Stop(data.mode).then(result => {
      if (result == "ok") {
        data.status = "off";
        data.btnStr = "Start!";
        changeLogo(img_neck_pains[Math.floor(Math.random() * img_neck_pains.length)])
      }
    });
  } else if (data.status == "off") {
    Start(data.mode).then(result => {
      if (result == "ok") {
        data.status = "on";
        data.btnStr = "Stop!";
        changeLogo(img_neck_pains_ok[Math.floor(Math.random() * img_neck_pains_ok.length)])
      }
    });
  }
}

function changeLogo(src) {
  console.log(src)
  data.logoImage = src
}

</script>

<template>
  <div id="workspace">
    <img id="logo" alt="Wails logo" :src="data.logoImage" />
    <div id="result" class="result">欢迎使用</div>
    <div id="input" class="input-box">
      <select class="input" id="group" v-model="data.mode">
        <option class="option-box" value="random">randmo</option>
        <option class="option-box" value="clockwise">clockwise</option>
        <option class="option-box" value="counterclockwise">counterclockwise</option>
      </select>
      <button class="btn" @click="btnfunc">{{ data.btnStr }}</button>
    </div>
  </div>
</template>

<style scoped>
.option-box{
  width: 100px;
  height: 35px;
}
#workspace {
  width: 100%;
  height: 100%;
  background-color: #333333;
}

#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
