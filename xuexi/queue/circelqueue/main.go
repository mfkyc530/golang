package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构提管理环形队列
type CircleQueue struct {
	maxSize int
	array [5]int
	head int  // 队首
	tail int  // 队尾
}

// 入队列
func (c *CircleQueue) PushQueue(val int) (err error){
	if c.IsFull(){
		return errors.New("queue full")
	}

	c.array[c.tail] = val
	c.tail = (c.tail + 1) % c.maxSize

	return
}

//出队列
func (c *CircleQueue) PopQueue()(val int, err error){
	if c.IsEmpty(){
		return 0, errors.New("queue empty")
	}

	// 取出
	val = c.array[c.head]
	c.head = (c.head + 1) % c.maxSize
	return val, err
}

// 显示队列
func (c *CircleQueue) ListQueue(){
	fmt.Println("环形队列如下：")
	//取出当前队列有多少个元素
	size := c.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	//设计一个辅助的变量，指向head
	tempHead := c.head
	for i := 0; i < size; i++{
		fmt.Printf("arr[%d] = %d \t", tempHead, c.array[tempHead])
		tempHead = (tempHead + 1) % c.maxSize
	}
	fmt.Println()
}

// 判断环形队列为满
func (c *CircleQueue) IsFull() bool{
	return (c.tail + 1) % c.maxSize == c.head
}

// 判断环形队列是空
func (c *CircleQueue) IsEmpty() bool{
	return c.tail == c.head
}

//取出环形队列有多少个元素
func (c *CircleQueue) Size() int{
	return (c.tail + c.maxSize - c.head) % c.maxSize
}

func main() {
	//初始化一个环形队列
	queue := & CircleQueue{
		maxSize:5,
		head:0,
		tail:0,

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
			err := queue.PushQueue(val)
			if err == nil {
				fmt.Println("加入队列ok")
			}else {
				fmt.Println(err.Error())
			}
		case "get":
			val, err := queue.PopQueue()
			if err != nil {
				fmt.Println(err.Error())
			}else {
				fmt.Println("从队列中取出了一个数 =", val)
			}

		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)

		}
	}
}
