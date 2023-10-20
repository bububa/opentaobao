package traderate

import (
	"sync"
)

// TopGetMixRateListParam 结构体
type TopGetMixRateListParam struct {
	// 筛选条件JSON
	TabFilter string `json:"tab_filter,omitempty" xml:"tab_filter,omitempty"`
	// 类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 酒店或商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// pageNo
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// pageSize
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolTopGetMixRateListParam = sync.Pool{
	New: func() any {
		return new(TopGetMixRateListParam)
	},
}

// GetTopGetMixRateListParam() 从对象池中获取TopGetMixRateListParam
func GetTopGetMixRateListParam() *TopGetMixRateListParam {
	return poolTopGetMixRateListParam.Get().(*TopGetMixRateListParam)
}

// ReleaseTopGetMixRateListParam 释放TopGetMixRateListParam
func ReleaseTopGetMixRateListParam(v *TopGetMixRateListParam) {
	v.TabFilter = ""
	v.BizType = 0
	v.ItemId = 0
	v.PageNo = 0
	v.PageSize = 0
	poolTopGetMixRateListParam.Put(v)
}
