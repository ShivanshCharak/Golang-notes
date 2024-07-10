
import "fmt"
func main(){
	myChannel := make(chan string)
	anotherChannel := make(anotherChannel string)
	go func add(){
		myChannel <- "Shivansh"
	}
	go func add2(){
		anotherChannel <- "Charak"
	}
	select{
	case anotherChannel := <- anotherChannel:
		fmt.Println(anotherChannel)
	case myChannel := <- myChannel:
		fmt.Println(msgChannel)
	}
}