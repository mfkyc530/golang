package sparsearray

type ValNode struct {
	Row int
	Col int
	Val int
}

func CreateSparsearray() []ValNode{
	var chessMap[11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 //白子

	var sparseArr []ValNode
	valNode := ValNode{
		Row:11,
		Col:11,
		Val:0,
	}

	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap{
		for j, v2 := range v{
			if v2 != 0{
				valNode := ValNode{
					Row:i,
					Col:j,
					Val:v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	return  sparseArr
}

func ReadSparsearray(data string){

}
