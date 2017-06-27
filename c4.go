package main

import (
  "fmt"
  "math"
  "strings"
  "io/ioutil"
  "encoding/hex"
)

func single_byte_xor(a []byte, b byte) []byte {
  result := make([]byte, len(a))
  for i, val := range a {
    result[i] = val ^ b
  }
  return result
}

func score_english_string(a []byte) float64 {
  english_freq := [28]float64{
    0.0651738, 0.0124248, 0.0217339, 0.0349835,  //'A', 'B', 'C', 'D',...
    0.1041442, 0.0197881, 0.0158610, 0.0492888,
    0.0558094, 0.0009033, 0.0050529, 0.0331490,
    0.0202124, 0.0564513, 0.0596302, 0.0137645,
    0.0008606, 0.0497563, 0.0515760, 0.0729357,
    0.0225134, 0.0082903, 0.0171272, 0.0013692,
    0.0145984, 0.0007836, 0.1918182, 0} //'Y', 'Z', ' '                  // V-Z
  length := len(a)
  var count [28]float64

  for _, c := range a {
    if (c >= 65 && c <= 90) {
      count[c - 65]++
      // uppercase A-Z
    } else if (c >= 97 && c <= 122) {
      count[c - 97]++
      // lowercase a-z
    } else if c == 32 {
      count[26]++
    } else if (c >= 33 && c <= 126) {
      count[27]++
      // numbers and punct.
    } else if (c == 9 || c == 10 || c == 13) {
      count[27]++
      // TAB, CR, LF
    } else {
      return math.Inf(1)
      // not printable ASCII = impossible(?)
    }
  }

  var chi2 float64
  for i, val := range count {
    expected := float64(length) * english_freq[i]
    difference := val - expected
    chi2 += difference * difference
  }

  return chi2
}

func main() {
  data, _ := ioutil.ReadFile("data/4.txt")
  output := "Now that the party is jumping\n"

  var bestPlaintext []byte
  var bestScore float64
  isSet := false

  for _, c := range strings.Split(string(data), "\n") {
    hex, _ := hex.DecodeString(c)
    for i := 0; i < 256; i++ {
      xored := single_byte_xor(hex, byte(i))
      score := score_english_string(xored)

      if !math.IsInf(score, 1) {
        if !isSet || score < bestScore {
          bestPlaintext = xored
          bestScore = score
          isSet = true
        }
      }
    }
  }

  if string(bestPlaintext) == output {
    fmt.Println("Challenge 4 passed!")
  } else {
    fmt.Println("Challenge 4 failed.")
  }
}
