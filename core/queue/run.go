package queue

import (
	"fmt"
)
func WrFunc (){
	arr :=make(chan int)
	defer close(arr)
	go func() {
		for i:=1;i<30 ;i++  {
			arr<-i
		}
	}()
	for{
		select {
		case <- arr:
			fmt.Println("222xxx")
		}
	}
}
