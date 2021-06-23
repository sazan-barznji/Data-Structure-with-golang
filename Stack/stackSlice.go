package main
import "fmt"

type stack struct {
	items []int
}

//adding elements to the stuck 
func (s *stack) push (value int ){
	s.items= append(s.items,value)
}

//remove value 
func (s *stack) pop() int {
	l:=len(s.items)-1
	toremove:=s.items[l] 
	s.items= s.items[:l] //start form the begining and leave the last 
	return toremove
}

func main(){
	mys := stack{}
	mys.push(100)
	mys.push(200)
	mys.push(300)
	mys.push(400)
	fmt.Println(mys)
	mys.pop()
	fmt.Println(mys)
	if ()
	mys.pop()
	fmt.Println(mys)


}