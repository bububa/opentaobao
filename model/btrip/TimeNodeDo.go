package btrip

import (
	"sync"
)

// TimeNodeDo 结构体
type TimeNodeDo struct {
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 时间类型
	TimeType string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	// 时间
	TimeStamp int64 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// 费用
	Cost int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// 费率
	CostPercent int64 `json:"cost_percent,omitempty" xml:"cost_percent,omitempty"`
}

var poolTimeNodeDo = sync.Pool{
	New: func() any {
		return new(TimeNodeDo)
	},
}

// GetTimeNodeDo() 从对象池中获取TimeNodeDo
func GetTimeNodeDo() *TimeNodeDo {
	return poolTimeNodeDo.Get().(*TimeNodeDo)
}

// ReleaseTimeNodeDo 释放TimeNodeDo
func ReleaseTimeNodeDo(v *TimeNodeDo) {
	v.Title = ""
	v.Content = ""
	v.TimeType = ""
	v.TimeStamp = 0
	v.Cost = 0
	v.CostPercent = 0
	poolTimeNodeDo.Put(v)
}
