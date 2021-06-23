package main
import "fmt"
type node struct{
	name string 
	grade int 
	next *node
}
var head *node
// adding new ndoe at the end of the linked list
//O(n)
func add (name string, grade int ){
	s:= node{name, grade, nil}
	if head ==nil{
		head = &s
	}else {
		current:=head 
		for current.next!=nil{
			current= current.next
		}
		current.next=&s
	}	
}
func addtoStart(name string, grade int){
	s:=node{name, grade, nil}
	s.next=head 
	head=&s
}

func print() {	
	c:=head
	for c!=nil{
		fmt.Println(c.name,c.grade)
		c=c.next
	}
}

func removeFirst(){
	head= head.next
}
func main(){
	add("sazan", 100)
	add("rezan", 100)
	//print()
	addtoStart ("fist",304)
	print()
	 removeFirst()
	 print()
}