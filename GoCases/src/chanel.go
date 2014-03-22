/*这个例子好像有问题哦*/
package main
import "fmt"
var counter int32 =1
func Count(ch chan int) {
	<-ch
    counter++
    fmt.Println("Counting", counter)
}
func main(){
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range(chs) {
		ch<-1
	}
}
