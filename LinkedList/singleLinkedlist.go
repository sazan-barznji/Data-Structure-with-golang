package main
import "fmt"
/////// with pointer 
type node struct {
	data int 
	name string
	next *node 
}
type linkedlist struct {
	head *node
	length int
}
//add int the begining 
func (l *linkedlist) add (n *node){
	second:=l.head
	l.head= n 
	l.head.next=second
	l.length++
}

func (l linkedlist)print(){
	print:=l.head
	for l.length!=0 {
		fmt.Printf("%d ",print.data,print.name)
		print= print.next
		l.length--
	}
	fmt.Printf("\n")
}
func (l *linkedlist) getdata(){
	var n int 
	var tempD int
	var nam string
	fmt.Println("enter how many data: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("enter data: ")
		fmt.Scan(&tempD)	
		fmt.Println("enter name: ")
		fmt.Scan(&nam)	
		l.add(&node{data:tempD,name:nam})
	}
}
func (l *linkedlist)delet(value int )  {
	if l.length==0{
		return
	}
	if l.head.data==value {
	l.head=l.head.next 
	l.length--
	return
	}
	prev:=l.head
	for prev.next.data!=value {
		if prev.next.next==nil{ //if this is the end 
			return
		}
		prev= prev.next
	}
	prev.next=prev.next.next
	l.length--	
}

func main(){
	list:=linkedlist{}
	list.getdata()
	list.print()
	list.delet(16)
	list.print()
}

//////// with value

// type node struct {
// 	data int 
// 	next *node 
// }
// type linkedlist struct {
// 	head *node
// 	length int
// }
// func (l *linkedlist)prepend(value int)  {
// 	newnode:= node{data:value}
// 	if l.head!=nil{
// 		newnode.next=l.head
// 		l.head=&newnode //moving head to new node
// 		l.length++
// 	}else{
// 		l.head=&newnode //if it the first node
// 		l.length++
// 	}
// 	return  
// }
// func (l *linkedlist) print(){
// 	if l.head==nil{
// 		return
// 	}
// 	current:=l.head
// 	for current!=nil{
// 		fmt.Println(current.data)
// 		current=current.next
// 	}
	
// }
// func (l *linkedlist) delete(value int)  {
// 	if l.length==0{   
// 		return
// 	}
// 	if l.head.data==value {
// 		l.head=l.head.next //set it to the next node
// 		l.length--
// 		return
// 	}
// 	//traversing the list 
// 	prev:=l.head 
// 	for prev.next.data != value{
// 		if prev.next.next==nil{ //we reached the tail 
// 			return
// 		}
// 		prev=prev.next
// 	}
// 	prev.next=prev.next.next
// 	l.length--	
// }

// func (l *linkedlist) getdata(){
// 	var n int 
// 	var tempD int
// 	fmt.Println("enter how many data: ")
// 	fmt.Scan(&n)

// 	for i := 0; i < n; i++ {
// 		fmt.Println("enter data: ")
// 		fmt.Scan(&tempD)	
// 		l.prepend(tempD)
// 	}
// }
// func main()  {
// 	list:=linkedlist{}
// 	list.getdata()
// 	list.print()
// }