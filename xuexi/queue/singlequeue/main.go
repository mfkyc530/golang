package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array [4]int // 数组
	front int // 表示指向队列首
	rear int // 表示指向队列尾部
}

// 添加队列
func(q *Queue) AddQueue (val int)(err error){
	// 判断队列是否已满
	if q.rear == q.maxSize -1 {
		return errors.New("queue full")
	}else {
		q.rear++ //rear 后移
		q.array[q.rear] = val
		return
	}
}

//显示队列   找到对首 ，然后到遍历到队尾
func (q *Queue) ShowQueue() {
	for i := q.front + 1; i <= q.rear; i++{
		fmt.Printf("array[%d] = %d \t", i, q.array[i])
	}

}

//从队列中取出数据
func(q *Queue) GetQueue()(val int, err error){
	if q.rear == q.front{
		return -1, errors.New("queue empty")
	}
	q.front++
	val = q.array[q.front]
	return val, err

}


func main() {
	// 先创建一个队列
	queue := Queue{
		maxSize: 5,
		front: -1,
		rear: -1,
	}

	var key string
	var val int

	for{
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err == nil {
				fmt.Println("加入队列ok")
			}else {
				fmt.Println(err.Error())
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			}else {
				fmt.Println("从队列中取出了一个数 =", val)
			}

		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)

		}
	}
}
