package tbtrade

import (
	"sync"
)

// OrderTopCursorQueryDto 结构体
type OrderTopCursorQueryDto struct {
	// 必填，查询开始时间，日期格式：yyyy-MM-dd HH:mm:ss.SSS
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 必填，查询结束时间（不可与开始日期跨天），日期格式：yyyy-MM-dd HH:mm:ss.SSS
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 页码游标值：上一次请求返回的游标id；第一次请求时不传
	CursorId string `json:"cursor_id,omitempty" xml:"cursor_id,omitempty"`
	// 1 udsmart. 2. 流量通
	ProdType int64 `json:"prod_type,omitempty" xml:"prod_type,omitempty"`
}

var poolOrderTopCursorQueryDto = sync.Pool{
	New: func() any {
		return new(OrderTopCursorQueryDto)
	},
}

// GetOrderTopCursorQueryDto() 从对象池中获取OrderTopCursorQueryDto
func GetOrderTopCursorQueryDto() *OrderTopCursorQueryDto {
	return poolOrderTopCursorQueryDto.Get().(*OrderTopCursorQueryDto)
}

// ReleaseOrderTopCursorQueryDto 释放OrderTopCursorQueryDto
func ReleaseOrderTopCursorQueryDto(v *OrderTopCursorQueryDto) {
	v.StartTime = ""
	v.EndTime = ""
	v.CursorId = ""
	v.ProdType = 0
	poolOrderTopCursorQueryDto.Put(v)
}
