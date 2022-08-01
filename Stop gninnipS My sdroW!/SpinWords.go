package kata
import (
  "strings"
  "bytes"
)

func SpinWords(str string) string {
  words := strings.Fields(str)
  var b bytes.Buffer
  for i, v := range words {
    // reverse words of length 5 or more
    if len(v) >= 5 {
      rns := []rune(v)
      for i, j := 0, len(rns)-1; i < j; i, j = i + 1, j - 1 {
        rns[i], rns[j] = rns[j], rns[i]
      }
      v = string(rns)
      b.WriteString(v)
    } else {
      b.WriteString(v)
    }
    if len(words) != i + 1{
      b.WriteString(" ")
    }
  }
  return b.String()
}
