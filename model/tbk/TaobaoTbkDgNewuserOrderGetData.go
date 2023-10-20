package tbk

import (
	"sync"
)

// TaobaoTbkDgNewuserOrderGetData 结构体
type TaobaoTbkDgNewuserOrderGetData struct {
	// resultList
	Results []TaobaoTbkDgNewuserOrderGetMapData `json:"results,omitempty" xml:"results>taobao_tbk_dg_newuser_order_get_map_data,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolTaobaoTbkDgNewuserOrderGetData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgNewuserOrderGetData)
	},
}

// GetTaobaoTbkDgNewuserOrderGetData() 从对象池中获取TaobaoTbkDgNewuserOrderGetData
func GetTaobaoTbkDgNewuserOrderGetData() *TaobaoTbkDgNewuserOrderGetData {
	return poolTaobaoTbkDgNewuserOrderGetData.Get().(*TaobaoTbkDgNewuserOrderGetData)
}

// ReleaseTaobaoTbkDgNewuserOrderGetData 释放TaobaoTbkDgNewuserOrderGetData
func ReleaseTaobaoTbkDgNewuserOrderGetData(v *TaobaoTbkDgNewuserOrderGetData) {
	v.Results = v.Results[:0]
	v.PageNo = 0
	v.PageSize = 0
	v.HasNext = false
	poolTaobaoTbkDgNewuserOrderGetData.Put(v)
}
