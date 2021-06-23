package main
import "fmt"

type node struct{
	value int 
	childern[2] *node
	parent *node
}
var root *node
func add (value int, target *node, dir int) *node{
	s:= node {value,nil,nil}
	if root !=nil{
		root =&s
		return &s
	}
	target.childern[dir]=&s
	s.parent=target 
	return &s

}
func remove (s *node){
	p:=s.parent
	if p!=nil{
		p=s.parent
	}
	for i := 0; i < len(p.childern); i++ {
		if p.childern[i]==s{
			p.childern[i]=nil
			break
		} 		  		 
	}
}
func print(){
	c:= root 
	for c!=nil {
		//get the children 
		left:=c.childern[0]
		right:=c.childern[1]

	}
}
func main(){
	s1:= add(100,nil,0)
	add (200,s1,0)
}