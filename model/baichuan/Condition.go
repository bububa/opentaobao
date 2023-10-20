package baichuan

import (
	"sync"
)

// Condition 结构体
type Condition struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 删除个数，必填，若超过最大值报错
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 对于删除，该字段无效
	StartId int64 `json:"start_id,omitempty" xml:"start_id,omitempty"`
	// 商品状态
	ItemStatus int64 `json:"item_status,omitempty" xml:"item_status,omitempty"`
}

var poolCondition = sync.Pool{
	New: func() any {
		return new(Condition)
	},
}

// GetCondition() 从对象池中获取Condition
func GetCondition() *Condition {
	return poolCondition.Get().(*Condition)
}

// ReleaseCondition 释放Condition
func ReleaseCondition(v *Condition) {
	v.StartTime = ""
	v.EndTime = ""
	v.Limit = 0
	v.StartId = 0
	v.ItemStatus = 0
	poolCondition.Put(v)
}
