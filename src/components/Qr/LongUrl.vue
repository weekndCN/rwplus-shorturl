<template>
  <v-textarea
    outlined
    dark
    auto-grow
    rows="3"
    v-model="lengthen"
    label="长URL"
    clearable
    append-outer-icon="mdi-send"
  ></v-textarea>
</template>

<script>
export default {
  props: {
    /* ✓ GOOD */
    // myProp: Number,
    // anotherProp: [Number, String],
    // myFieldWithBadType: {
    //   type: Object,
    //   default: function() {
    //     return {}
    //   },
    // },
    // myOtherFieldWithBadType: {
    //   type: Number,
    //   default: 1,
    // },
    lengthen: {
      type: String,
      value: ""
    }
  }
};
</script>

<style>
</style>