package main
import "fmt"
type node struct {
	data int 
	next *node 
}
type linkedlist struct {
	head *node
	length int
}
func (l *linkedlist) add (n *node){
	second:=l.head
	l.head= n 
	l.head.next=second
	l.length++
}
func (l linkedlist)print(){
	print:=l.head
	for l.length!=0 {
		fmt.Printf("%d ",print.data)
		print= print.next
		l.length--
	}
	fmt.Printf("\n")
}
func (l *linkedlist) getdata(){
	var n int 
	var tempD int
	fmt.Println("enter how many data: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("enter data: ")
		fmt.Scan(&tempD)	
		l.add(&node{data:tempD})
	}
}
func main(){
	list:=linkedlist{}
	list.getdata()
	list.print()
}