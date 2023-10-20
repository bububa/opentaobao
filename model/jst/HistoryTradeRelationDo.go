package jst

import (
	"sync"
)

// HistoryTradeRelationDo 结构体
type HistoryTradeRelationDo struct {
	// 标签名称
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 标签值，json格式
	TagValue string `json:"tag_value,omitempty" xml:"tag_value,omitempty"`
	// 记录的创建时间
	GmtCreated string `json:"gmt_created,omitempty" xml:"gmt_created,omitempty"`
	// 记录的最新修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 订单标签记录id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 订单id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 标签类型       1：官方标签      2：自定义标签
	TagType int64 `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
	// 该标签在消费者端是否显示,0:不显示,1：显示
	Visible int64 `json:"visible,omitempty" xml:"visible,omitempty"`
}

var poolHistoryTradeRelationDo = sync.Pool{
	New: func() any {
		return new(HistoryTradeRelationDo)
	},
}

// GetHistoryTradeRelationDo() 从对象池中获取HistoryTradeRelationDo
func GetHistoryTradeRelationDo() *HistoryTradeRelationDo {
	return poolHistoryTradeRelationDo.Get().(*HistoryTradeRelationDo)
}

// ReleaseHistoryTradeRelationDo 释放HistoryTradeRelationDo
func ReleaseHistoryTradeRelationDo(v *HistoryTradeRelationDo) {
	v.TagName = ""
	v.TagValue = ""
	v.GmtCreated = ""
	v.GmtModified = ""
	v.Id = 0
	v.Tid = 0
	v.TagType = 0
	v.Visible = 0
	poolHistoryTradeRelationDo.Put(v)
}
