package main
import "fmt"

type node struct{
	value int 
	right *node
	left *node
	parent *node
}
var root *node
func add (value int, target *node, dir int) *node{
	s:= node {value,nil,nil,nil}

	if root !=nil{
		root =&s
		return &s
	}
	if dir==0 {
		//add to left 
		target.left= &s 
	}else {
		//add to right 
		target.right= &s
	}
	s.parent=target 
	return &s



}
func main(){
	s1:= add(100,nil,0)
	add (200,s1,0)
}