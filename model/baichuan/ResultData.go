package baichuan

import (
	"sync"
)

// ResultData 结构体
type ResultData struct {
	// 商品id列表
	ItemList []int64 `json:"item_list,omitempty" xml:"item_list>int64,omitempty"`
	// 具体内容
	DataList []int64 `json:"data_list,omitempty" xml:"data_list>int64,omitempty"`
	// 商品数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolResultData = sync.Pool{
	New: func() any {
		return new(ResultData)
	},
}

// GetResultData() 从对象池中获取ResultData
func GetResultData() *ResultData {
	return poolResultData.Get().(*ResultData)
}

// ReleaseResultData 释放ResultData
func ReleaseResultData(v *ResultData) {
	v.ItemList = v.ItemList[:0]
	v.DataList = v.DataList[:0]
	v.Count = 0
	poolResultData.Put(v)
}
