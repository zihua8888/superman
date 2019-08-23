package main

import "fmt"

var complete chan int = make(chan int)
var ch = make(chan int,3)
func loop(){
	for i := 0;i< 10;i++{
		fmt.Printf("%d",i)
	}


	complete <- 0
}

func hello()  {
	ch <- 1
}

func foo(id int)  {
	ch <- id
}

func main() {
	go loop()
	go hello()

	<- complete //直到线程跑完，渠道消息.main在此阻塞住

	<- ch

	//开启5个routine
	for i := 0;i<5;i++{
		go foo(i)
	}
	//取出信道中的数据
	for i := 0; i <5;i++{
		fmt.Println(<- ch)
	}


	ch <- 11
	ch <- 22
	ch <- 33
	close(ch)
	for v := range ch {
		fmt.Println(v)
		//if len(ch) <= 0{//如果现有数据量为0，跳出循环
		//	break
		//}
	}

}
