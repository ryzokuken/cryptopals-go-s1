package main

import (
  "fmt"
  "encoding/hex"
)

func repeatingKeyXor(plaintext, key []byte) []byte {
  ciphertext := make([]byte, len(plaintext))
  length := len(key)
  for i, val := range plaintext {
    ciphertext[i] = val ^ key[i % length]
  }
  return ciphertext
}

func main() {
  input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
  key := "ICE"
  output := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

  // hexed := make([]byte, hex.EncodedLen(len([]byte(input))))
  // hex.Encode(hexed, input)

  if hex.EncodeToString(repeatingKeyXor([]byte(input), []byte(key))) == output {
    fmt.Println("Challenge 5 passed!")
  } else {
    fmt.Println("Challenge 5 failed.")
  }
}
