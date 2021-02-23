package main
import "fmt"

//receiver结构体
type receiver struct{
	name string
}
//空接口
type message interface{
}
//receive方法
func (r receiver)receive(v message){
	//switch分支
	switch v.(type){
	case int: fmt.Println("这个是int")
	case bool:fmt.Println("这个是bool")
	case string:fmt.Println("这个是string")
	}
}


var r receiver
var m message
func main(){
	r := receiver{"r"}
	r.receive("你好")
	r.receive(32)
	r.receive(true)
}