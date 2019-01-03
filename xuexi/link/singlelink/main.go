package main

import "fmt"

type HeroNode struct {
	no int
	name string
	nickname string
	next *HeroNode
}

//给链表插入一个节点  编写第一种插入方法，在单链表的最后加入
func InsertHeroNode(h *HeroNode, newHeroNode *HeroNode){
	//1先找到该链表最后这个节点
	//2创建一个辅助节点
	temp := h
	for {
		if temp.next == nil{  //表示找到最后
			break
		}
		temp = temp.next //让temp不断指向下一个节点
	}

	// 3将newHeroNode加入到链表最后
	temp.next = newHeroNode
}

//给链表插入一个节点  编写第er种插入方法，根据no的编号从小到大插入
func InsertHeroNode2(h *HeroNode, newHeroNode *HeroNode){
	//1找到适当的结点
	//2创建一个辅助节点
	temp := h

	flag := true
	//让插入的结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil{
			break
		}else if temp.next.no > newHeroNode.no{
			//说明newheronode 就应该插入到temp后面
			break
		}else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next

	}
	if !flag {
		fmt.Println("对不起，已经存在no =", newHeroNode.no)
	}else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}

}

//显示链表的所有节点信息
func ListHeroNode(h *HeroNode){
	//1.创建一个辅助节点
	temp := h
	if temp.next == nil {
		fmt.Printf("链表为空。。。。")
		return
	}

	//遍历这个链表
	for {
		fmt.Printf("[%d, %s, %s] ==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil{
			break
		}
	}
}

func main() {
	// 1先创建一个头节点
	head := &HeroNode{

	}

	// 2创建一个新的HeroNode
	hero1 := &HeroNode{
		no:1,
		name:"宋江",
		nickname:"及时雨",

	}
	hero2 := &HeroNode{
		no:2,
		name:"卢俊义",
		nickname:"玉麒麟",

	}
	hero3 := &HeroNode{
		no:3,
		name:"林冲",
		nickname:"豹子头",

	}

	//加入
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)
	//显示
	ListHeroNode(head)
}