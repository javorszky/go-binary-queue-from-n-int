package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func binaryQueueConvert(num int64) *list.List {
	l := list.New()
	for i := int64(0); i <= num; i++  {
		l.PushBack(strconv.FormatInt(i, 2))
	}

	return l
}

func binaryQueueDisplay(num int) *list.List {
	l := list.New()
	for i := int(0); i <= num; i++  {
		l.PushBack(i)
	}

	return l
}


func main() {
	queue := binaryQueueConvert(20)
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value);
	}

	fmt.Println("Displaying the next version")

	queue2 := binaryQueueDisplay(20)
	for f := queue2.Front(); f != nil; f = f.Next() {
		fmt.Println(fmt.Sprintf("%b", f.Value))
	}
}
