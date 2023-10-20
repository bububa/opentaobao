package tuanhotel

import (
	"sync"
)

// TopSkuDailyInfo 结构体
type TopSkuDailyInfo struct {
	// 日期
	D string `json:"d,omitempty" xml:"d,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
}

var poolTopSkuDailyInfo = sync.Pool{
	New: func() any {
		return new(TopSkuDailyInfo)
	},
}

// GetTopSkuDailyInfo() 从对象池中获取TopSkuDailyInfo
func GetTopSkuDailyInfo() *TopSkuDailyInfo {
	return poolTopSkuDailyInfo.Get().(*TopSkuDailyInfo)
}

// ReleaseTopSkuDailyInfo 释放TopSkuDailyInfo
func ReleaseTopSkuDailyInfo(v *TopSkuDailyInfo) {
	v.D = ""
	v.Price = ""
	v.Stock = 0
	poolTopSkuDailyInfo.Put(v)
}
