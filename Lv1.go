package main
import "fmt"


type person struct{
	name string
	age int
}
type dove interface{
	gugugu()
}
type repeater interface{
	repeate(string)
}
type sourgrape interface{
	sour()
}
func (p person) gugugu(){
	fmt.Println("又鸽了")
}
func (p person) repeate(s string){
	fmt.Println(s)
}
func (p person) sour(){
	fmt.Println("sour")
}

var p person



func main(){
	p := person{"lc",19}
	var g dove
	g = p
	g.gugugu()
	var r repeater
	r = p
	r.repeate("I am repeater")
	var lemon sourgrape
	lemon = p
	lemon.sour()
}