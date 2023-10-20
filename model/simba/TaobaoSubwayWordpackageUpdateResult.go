package simba

import (
	"sync"
)

// TaobaoSubwayWordpackageUpdateResult 结构体
type TaobaoSubwayWordpackageUpdateResult struct {
	// 更新数目
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

var poolTaobaoSubwayWordpackageUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayWordpackageUpdateResult)
	},
}

// GetTaobaoSubwayWordpackageUpdateResult() 从对象池中获取TaobaoSubwayWordpackageUpdateResult
func GetTaobaoSubwayWordpackageUpdateResult() *TaobaoSubwayWordpackageUpdateResult {
	return poolTaobaoSubwayWordpackageUpdateResult.Get().(*TaobaoSubwayWordpackageUpdateResult)
}

// ReleaseTaobaoSubwayWordpackageUpdateResult 释放TaobaoSubwayWordpackageUpdateResult
func ReleaseTaobaoSubwayWordpackageUpdateResult(v *TaobaoSubwayWordpackageUpdateResult) {
	v.Result = 0
	poolTaobaoSubwayWordpackageUpdateResult.Put(v)
}
