package koubeimall

import (
	"sync"
)

// ItemGroupContentDetail 结构体
type ItemGroupContentDetail struct {
	// 金额
	ContentAmount string `json:"content_amount,omitempty" xml:"content_amount,omitempty"`
	// 单位，份，杯
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 规格
	Spec string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 名称，eg：红酒
	ContentName string `json:"content_name,omitempty" xml:"content_name,omitempty"`
	// 份数
	ContentCount int64 `json:"content_count,omitempty" xml:"content_count,omitempty"`
}

var poolItemGroupContentDetail = sync.Pool{
	New: func() any {
		return new(ItemGroupContentDetail)
	},
}

// GetItemGroupContentDetail() 从对象池中获取ItemGroupContentDetail
func GetItemGroupContentDetail() *ItemGroupContentDetail {
	return poolItemGroupContentDetail.Get().(*ItemGroupContentDetail)
}

// ReleaseItemGroupContentDetail 释放ItemGroupContentDetail
func ReleaseItemGroupContentDetail(v *ItemGroupContentDetail) {
	v.ContentAmount = ""
	v.Unit = ""
	v.Spec = ""
	v.ContentName = ""
	v.ContentCount = 0
	poolItemGroupContentDetail.Put(v)
}
