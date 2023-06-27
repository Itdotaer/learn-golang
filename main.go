package main

//
//import "fmt"
//
//func main() {
//	s := make([]string, 0)
//	p := make([]Plusser, 0)
//
//	s = append(s, "A")
//	p = append(p, PI{msg: "B"})
//	s = append(s, "C")
//	p = append(p, PI{msg: "D"})
//
//	ss := SI{"jioj"}
//	sss := ss.(Stringer).String()
//
//	rs := ConcatTo(s, p)
//	println(rs)
//}
//
//func Print[T any](t T) {
//	fmt.Printf("printing type: %T\n", t)
//}
//
//type Stringer interface {
//	String() string
//}
//
//type SI struct {
//	msg string
//}
//
//func (s SI) String() string {
//	return s.msg
//}
//
//type Plusser interface {
//	Plus(string) string
//}
//
//type PI struct {
//	msg string
//}
//
//func (p PI) Plus(in string) string {
//	return p.msg + "--" + in
//}
//
//func ConcatTo[S Stringer, P Plusser](s []S, p []P) string {
//	r := ""
//	for i, v := range s {
//		r += p[i].Plus(v.String())
//	}
//	return r
//}
