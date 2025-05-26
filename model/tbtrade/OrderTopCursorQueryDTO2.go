package tbtrade

import (
	"sync"
)

// OrderTopCursorQueryDTO2 结构体
type OrderTopCursorQueryDTO2 struct {
	// 必填，查询开始时间，日期格式：yyyy-MM-dd HH:mm:ss.SSS
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 必填，查询结束时间（不可与开始日期跨天），日期格式：yyyy-MM-dd HH:mm:ss.SSS
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 页码游标值：上一次请求返回的游标id；第一次请求时不传
	CursorId string `json:"cursor_id,omitempty" xml:"cursor_id,omitempty"`
	// 1 udsmart. 2. 流量通
	ProdType int64 `json:"prod_type,omitempty" xml:"prod_type,omitempty"`
}

var poolOrderTopCursorQueryDTO2 = sync.Pool{
	New: func() any {
		return new(OrderTopCursorQueryDTO2)
	},
}

// GetOrderTopCursorQueryDTO2() 从对象池中获取OrderTopCursorQueryDTO2
func GetOrderTopCursorQueryDTO2() *OrderTopCursorQueryDTO2 {
	return poolOrderTopCursorQueryDTO2.Get().(*OrderTopCursorQueryDTO2)
}

// ReleaseOrderTopCursorQueryDTO2 释放OrderTopCursorQueryDTO2
func ReleaseOrderTopCursorQueryDTO2(v *OrderTopCursorQueryDTO2) {
	v.StartTime = ""
	v.EndTime = ""
	v.CursorId = ""
	v.ProdType = 0
	poolOrderTopCursorQueryDTO2.Put(v)
}
