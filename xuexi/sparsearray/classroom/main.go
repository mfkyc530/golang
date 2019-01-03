package classroom

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main(){
	// 创建一个原始数组
	var chessMap[11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 白子

	// 输出看看原始的数组
	for _, v := range chessMap{
		for _,v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 转成稀疏数组  想->算法
	// 思路 1遍历 chessMap，如果我们发现又一个元素的值不为0，创建一个node结构体
	//       将其放入到对应的切片中

	var sparseArr []ValNode

	// 标准的稀疏数组应该还有一个表示记录原始的二维数组的规模（行和列、默认值）
	valNode := ValNode{
		row:11,
		col:11,
		val:0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap{
		for j, v2 := range v {
			if v2 != 0{
				// 创建一个ValNode 值节点
				valNode := ValNode{
					row:i,
					col:j,
					val:v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	fmt.Printf("当前的稀疏数组是\n")
	for i, valNode := range sparseArr{
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}


	// 将这个稀疏数组，存盘

	// 如何恢复原始的数组   1。打开存盘文件 恢复原始数组
	var chessMap2 [11][11]int
	for i, valNode := range sparseArr{
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}

	}

	fmt.Printf("恢复后的原始数据\n")
	for _, v := range chessMap2{
		for _,v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
