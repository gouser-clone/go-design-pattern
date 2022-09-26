package simplefactory

import "fmt"

type SayApi interface{
  Say(name string) string
}

type Hello struct{}
type Hi struct{}
func NewSApi(t int) SayApi {
      if t == 1 {
      return &Hello{}
      }else if t == 2 {
      return &Hi{}
      }    
    return nil
}

func(he *Hello)Say(name string) string{
     return fmt.Sprintf("Hello, %s", name)
}
func(h *Hi)Say(name string) string{
     return fmt.Sprintf("Hi, %s", name)
}
func main(){

// he := &Hello{}
// rst := he.Say("tom")
// fmt.Println(rst)
// h := &Hi{}
// h.Say("jerry")


he := NewSApi(1)
var rst = he.Say("tom")
 fmt.Println(rst)
 fmt.Println()
h := NewSApi(2)
var rs = h.Say("jerry")
 fmt.Println(rs)


}
