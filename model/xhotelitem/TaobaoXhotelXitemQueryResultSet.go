package xhotelitem

import (
	"sync"
)

// TaobaoXhotelXitemQueryResultSet 结构体
type TaobaoXhotelXitemQueryResultSet struct {
	// 查询到的 x 元素
	XItems []HotelXitemDo `json:"x_items,omitempty" xml:"x_items>hotel_xitem_do,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 记录总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolTaobaoXhotelXitemQueryResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelXitemQueryResultSet)
	},
}

// GetTaobaoXhotelXitemQueryResultSet() 从对象池中获取TaobaoXhotelXitemQueryResultSet
func GetTaobaoXhotelXitemQueryResultSet() *TaobaoXhotelXitemQueryResultSet {
	return poolTaobaoXhotelXitemQueryResultSet.Get().(*TaobaoXhotelXitemQueryResultSet)
}

// ReleaseTaobaoXhotelXitemQueryResultSet 释放TaobaoXhotelXitemQueryResultSet
func ReleaseTaobaoXhotelXitemQueryResultSet(v *TaobaoXhotelXitemQueryResultSet) {
	v.XItems = v.XItems[:0]
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.TotalCount = 0
	poolTaobaoXhotelXitemQueryResultSet.Put(v)
}
