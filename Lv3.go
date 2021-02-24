package main
import "fmt"

//Author结构体
type Author struct{
	Name string		//名字
	VIP bool		//是否是会员
	Focus int		//关注人数
}

//bilibili接口
type bilibili interface{
	zan()
	collect()
	toubi()
	three()
}
//Vedio结构体
type Vedio struct{
	AuthorName string
	VedioName string
}

func (v Vedio)zan(){
	fmt.Println("作品太好，必须点赞")
}
func (v Vedio)collect(){
	fmt.Println("啥也不说了，收藏")
}
func (v Vedio)toubi(){
	fmt.Println("币有你的")
}
func (v Vedio)three(){
	fmt.Println("一键三连")
}

//pubilic(发布视频)方法
func (author Author)public(an string,vn string)(v Vedio){
	v.AuthorName = an
	v.VedioName = vn
	return v	//返回Vedio结构体
}


var author Author
var vedio Vedio
var   bili   bilibili
func main(){
	author := Author{"lc",true,99}
	vedio = author.public("lc","one piece")
	bili = vedio
	bili.zan()
	bili.collect()
	bili.toubi()
	bili.three()
}