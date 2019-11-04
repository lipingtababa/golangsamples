package main
import "fmt"
import "time"

func main(){
	pipeline()
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}

func pipeline(){
	in :=make(chan int)	//Make a channel for producer
	out := make(chan int)	//Make a channel for consumer
	go producer(in)		//Produce an infinite list of integer
	for i := 0; i<10;i++{
		go worker(in, out)	//Transition the list
	}
	go consumer(out, 100)

	fmt.Println("pipeline done")
}

func factor(i int) int{
	return  i
}

func worker (in chan int, out chan int){
	for {
		order := <-in
		result := factor(order)
		out <- result
	}
}

/*
func worker2 (in chan int, out chan int){
	for i :=0; i<10; i++{
		order := <-in
		result := factor(order)
		out <- result
	}
}
*/

func producer (out chan int ){
	for order :=0; ; order++ {
		out <- order
	}
}

func consumer(in chan int, n int){
	for i :=0; i<n; i++ {
		result := <-in
		fmt.Println("Consumed", result)
	}
}
