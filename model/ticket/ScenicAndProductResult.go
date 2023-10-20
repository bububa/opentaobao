package ticket

import (
	"sync"
)

// ScenicAndProductResult 结构体
type ScenicAndProductResult struct {
	// 景点列表
	ScenicList []Scenic `json:"scenic_list,omitempty" xml:"scenic_list>scenic,omitempty"`
}

var poolScenicAndProductResult = sync.Pool{
	New: func() any {
		return new(ScenicAndProductResult)
	},
}

// GetScenicAndProductResult() 从对象池中获取ScenicAndProductResult
func GetScenicAndProductResult() *ScenicAndProductResult {
	return poolScenicAndProductResult.Get().(*ScenicAndProductResult)
}

// ReleaseScenicAndProductResult 释放ScenicAndProductResult
func ReleaseScenicAndProductResult(v *ScenicAndProductResult) {
	v.ScenicList = v.ScenicList[:0]
	poolScenicAndProductResult.Put(v)
}
