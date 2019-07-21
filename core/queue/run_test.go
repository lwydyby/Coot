package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
		ch:= make(chan string)
		go func() {
			ch<-"111"
		}()
	for  {
		select {
		case c:=<-ch:
			fmt.Println("ddddd",c)
		default:
		}
	}
	time.Sleep(time.Second*10)
}