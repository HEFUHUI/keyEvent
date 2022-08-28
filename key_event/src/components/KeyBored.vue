<template>
  <div style="flex: 1">
    <div id="view" style="width: 100%"></div>
  </div>
</template>
<script>

export default {
  props: {
    data:{
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      keyData: [],
      keyLayout:{
        "q": "1",
        "w": "2",
        "e": "3",
        "r": "4",
        "t": "5",
        "y": "6",
        "u": "7",
        "i": "8",
        "o": "9",
        "p": "0",
        "a": "q",
        "s": "w",
        "d": "e",
        "f": "r",
        "g": "t",
        "h": "y",
        "j": "u",
        "k": "i",
        "l": "o",
        "z": "p",
        "x": "a",
        "c": "s",
        "v": "d",
        "b": "f",
        "n": "g",
        "m": "h",
        ",": "j",
        ".": "k",
        "/": "l",
        " ": "z",
        ";": "x",
        "'": "c",
        "[": "v",
        "]": "b",
        "\\": "n",
        "=": "m",
        "`": ",",
        "~": ".",
        "!": "/",
        "@": " ",
        "#": ";",
        "$": "'",
        "%": "[",
        "^": "]",
        "&": "\\",
        "*": "=",
        "(": "`",
        ")": "~",
        "_": "!",
        "+": "@",
        "-": "#",
        "|": "$",
        "{": "%",
        "}": "^",
        ":": "&",
        "\"": "*",
        "<": "(",
        ">": ")",
        "?": "_",
      }
    };
  },
  mounted() {
    this.renderView();
  },
  methods:{
    renderView(){
      
    }
  },
};
</script>
<style scoped>
div{
  transition: .2s;
}

</style>