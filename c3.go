package main

import (
  "fmt"
  "math"
  "encoding/hex"
)

func c3() {
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
