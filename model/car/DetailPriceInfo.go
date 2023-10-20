package car

import (
	"sync"
)

// DetailPriceInfo 结构体
type DetailPriceInfo struct {
	// 费用金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 费用类型名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 费用类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolDetailPriceInfo = sync.Pool{
	New: func() any {
		return new(DetailPriceInfo)
	},
}

// GetDetailPriceInfo() 从对象池中获取DetailPriceInfo
func GetDetailPriceInfo() *DetailPriceInfo {
	return poolDetailPriceInfo.Get().(*DetailPriceInfo)
}

// ReleaseDetailPriceInfo 释放DetailPriceInfo
func ReleaseDetailPriceInfo(v *DetailPriceInfo) {
	v.Amount = ""
	v.Name = ""
	v.Type = ""
	poolDetailPriceInfo.Put(v)
}
