package main

import (
  "fmt"
  "math"
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
  english_freq := [26]float64{
    0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015,  // A-G
    0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749,  // H-N
    0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758,  // O-U
    0.00978, 0.02360, 0.00150, 0.01974, 0.00074 }                   // V-Z
  length := len(a)
  var count [26]float64

  for _, c := range a {
    if (c >= 65 && c <= 90) {
      count[c - 65]++;        // uppercase A-Z
    } else if (c >= 97 && c <= 122) {
      count[c - 97]++;  // lowercase a-z
    } else if (c >= 32 && c <= 126) {
      length--;        // numbers and punct.
    } else if (c == 9 || c == 10 || c == 13) {
      length--;  // TAB, CR, LF
    } else {
      // fmt.Println("rejecting for:", c)
      return math.Inf(1);  // not printable ASCII = impossible(?)
    }
  }

  var chi2 float64
  for i, val := range count {
    expected := float64(length) * english_freq[i]
    difference := val - expected
    chi2 += difference * difference / expected
  }

  return chi2
}

func main() {
  input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  output := "Cooking MC's like a pound of bacon"
  hex, _ := hex.DecodeString(input)

  var minKey byte
  var minScore float64
  isSet := false
  for i := 0; i < 256; i++ {
    xored := single_byte_xor(hex, byte(i))
    score := score_english_string(xored)

    if !math.IsInf(score, 1) {
      if !isSet || score < minScore {
        minKey = byte(i)
        minScore = score
        isSet = true
      }
    }
  }

  if string(single_byte_xor(hex, minKey)) == output {
    fmt.Println("Challenge 3 passed!")
  } else {
    fmt.Println("Challenge 3 failed.")
  }
}
