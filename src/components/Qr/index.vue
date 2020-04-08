<template>
  <v-container
    fill-height
    fluid
    style="background: linear-gradient(to right, #24243e, #141e30, #0f0c29);"
  >
    <v-row align="center" justify="center">
      <v-col cols="12" sm="6" md="4" xs="4">
        <div class="title font-weight-bold text-center grey--text mb-7">RWPLUS SHORTURL</div>
        <v-textarea
          outlined
          dark
          auto-grow
          rows="3"
          v-model="lengthen"
          label="长URL"
          clearable
          append-outer-icon="mdi-send"
          @keydown.enter.prevent="getshorten"
        ></v-textarea>
        <v-text-field
          dark
          dense
          hide-details
          color="transparent"
          outlined
          label="短URL"
          v-model="shorten"
          append-outer-icon="mdi-arrange-bring-forward"
        ></v-text-field>
        <Qr />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Qr from "./Qr";

export default {
  name: "qrLayout",
  data: () => ({
    lengthen: "http://www.rwplus.cn/LYVqxgZ?editable=true&editors=101",
    shorten: "http://t.rwplus.cn/abcde",
    qrcode: ""
  }),
  methods: {
    clearMessage() {
      this.Lengthen = "http://www.rwplus.cn/LYVqxgZ?editable=true&editors=101";
      this.shorten = "http://t.rwplus.cn/abcde";
    },
    getshorten() {
      let data = new FormData();
      data.append("lengthen", this.lengthen);
      this.service
        .post("shorten", data, {
          headers: {
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Credentials": "true",
            "Content-Type": "application/x-www-form-urlencoded;charset=utf-8"
          },
          withCredentials: true
        })
        .then(resp => {
          console.log(resp.data.data.shorten);
          this.shorten = "http://webcode/j/" + resp.data.data.shorten;
          this.qrcode = "http://webcode/qrcode/" + resp.data.data.qrcode;
          console.log(this.shorten, this.qrcode);
        })
        .catch(err => {
          console.log("请求失败 " + err);
        });
    }
  },
  components: {
    Qr
  }
};
</script>

<style>
</style>