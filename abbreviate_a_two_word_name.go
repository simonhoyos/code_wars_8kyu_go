package kata

import "strings"

func AbbrevName(name string) string{
  s := strings.Split(name, " ")

  var res []string
  for _, r := range s {
    res = append(res, strings.ToUpper(string(r[0])))
  }

  return strings.Join(res, ".")
}

