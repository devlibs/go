package main
import(
	"fmt"
)

type human struct {
	Sex int
}

type teacher struct{
	human
	Sex int
	Age int
	Name string
}

type student struct{
	Age int
	Name string
}



func main() {
	
	a:=teacher{Name:"liuqing",Age:30,human:human{Sex:11}}
	a.Sex=10
	a.human.Sex=22

	fmt.Println(a)

}
