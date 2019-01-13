package BLC

import (
	"fmt"
	"math/big"
)

const targetBites  = 24
//	创教pow结构体
type ProofOfWork struct {
//	当前需要验证的区块
	block *Block
//	大数存储
	target *big.Int
}
// 创建newProofOfWork的方法

func newProofOfWork(block *Block) *ProofOfWork {
//	输出target
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBites))
	fmt.Println(target)
	fmt.Println("-----------")
	pow := &ProofOfWork{block,target}

	return pow
}

func (pow *ProofOfWork)Run()  {
	
}