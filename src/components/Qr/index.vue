<template>
  <v-container
    fill-height
    fluid
    style="background: linear-gradient(to right, #24243e, #141e30, #0f0c29);"
  >
    <v-row align="center" justify="center">
      <v-col cols="12" sm="6" md="4" xs="4">
        <div class="title font-weight-bold text-center grey--text mb-7">RWPLUS SHORTURL</div>
        <LongUrl :lengthen="lengthen" />
        <ShortUrl :shorten="shorten" />
        <Qr />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import LongUrl from "./LongUrl";
import ShortUrl from "./ShortUrl";
import Qr from "./Qr";

export default {
  name: "qrLayout",
  data: () => ({
    lengthen: "http://www.rwplus.cn/LYVqxgZ?editable=true&editors=101",
    shorten: "http://t.rwplus.cn/abcde"
  }),
  methods: {
    // sendMessage() {
    //   this.clearMessage();
    // },
    // clearMessage() {
    //   this.Lengthen = "http://www.rwplus.cn/LYVqxgZ?editable=true&editors=101";
    //   this.shorten = "http://t.rwplus.cn/abcde";
    // }
  },
  components: {
    LongUrl,
    ShortUrl,
    Qr
  }
};
</script>

<style>
</style>