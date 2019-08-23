package main

import "fmt"

func main() {
	//who := "wor"
	//if len(os.Args) >1{
	//	who = strings.Join(os.Args[1:],"")
	//}
	//
	//fmt.Println("hello",who)
	//fmt.Println(os.Args[0:])
	//fmt.Println(os.Args[1:])
	//fmt.Println(len(os.Args))
	//
	////a := "hahaha"
	////b := "hehehe"
	////c := strings.Join([]string{a,b},",")
	////println(c)
	//
	//var a = [][]string{{"1"},{"2"}}
	//b := [][]string{{"3"},{"4"}}
	//fmt.Println(a,b)

	//createCounter(3)
	//fmt.Println(createCounter(3))
	//fmt.Printf("%d",createCounter(3))

	//aa := make(chan int)
	//aa <- 2
	//bb,ok :=<- aa
	//if ok {
	//	fmt.Println(bb)
	//}

	var messages chan string = make(chan string)
	go func(message,m string) {

		messages <- message //存消息
		messages <- m
	}("ping","jjj")

	fmt.Println(<-messages,) //取消息
	fmt.Println(<-messages,) //取消息
}


func createCounter(start int) chan int{
	next := make(chan int)
	go func(i int) {
		next <- i
		i++
	}(start)
	return next

}

