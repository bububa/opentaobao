package icbu

import (
	"sync"
)

// ProductScoreInfoResult 结构体
type ProductScoreInfoResult struct {
	// 质量分
	FinalScore string `json:"final_score,omitempty" xml:"final_score,omitempty"`
	// 精品标，，返回字段中 boutique_tag 含义： 1 精品 2 普通品 3 低质品 4 实力优品
	BoutiqueTag int64 `json:"boutique_tag,omitempty" xml:"boutique_tag,omitempty"`
}

var poolProductScoreInfoResult = sync.Pool{
	New: func() any {
		return new(ProductScoreInfoResult)
	},
}

// GetProductScoreInfoResult() 从对象池中获取ProductScoreInfoResult
func GetProductScoreInfoResult() *ProductScoreInfoResult {
	return poolProductScoreInfoResult.Get().(*ProductScoreInfoResult)
}

// ReleaseProductScoreInfoResult 释放ProductScoreInfoResult
func ReleaseProductScoreInfoResult(v *ProductScoreInfoResult) {
	v.FinalScore = ""
	v.BoutiqueTag = 0
	poolProductScoreInfoResult.Put(v)
}
