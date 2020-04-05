<template>
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
</template>

<script>
export default {
  props: {
    shorten: {
      type: String,
      value: ""
    }
  }
};
</script>

<style>
</style>