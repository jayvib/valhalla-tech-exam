package convert

import (
  "strconv"
  "strings"
)

func StringToArrayOfInt(in string) ([]int, error) {
  out := make([]int, 0)
  splits := strings.Split(in, " ")

  for _, s := range splits {
    // convert string of int to int type
    i, err := strconv.Atoi(s)
    if err != nil {
      return nil, err
    }
    out = append(out, i)
  }
  return out, nil
}
