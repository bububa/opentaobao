package caipiao

import (
	"sync"
)

// LotteryType 结构体
type LotteryType struct {
	// 彩种名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 彩种ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolLotteryType = sync.Pool{
	New: func() any {
		return new(LotteryType)
	},
}

// GetLotteryType() 从对象池中获取LotteryType
func GetLotteryType() *LotteryType {
	return poolLotteryType.Get().(*LotteryType)
}

// ReleaseLotteryType 释放LotteryType
func ReleaseLotteryType(v *LotteryType) {
	v.Name = ""
	v.Id = 0
	poolLotteryType.Put(v)
}
