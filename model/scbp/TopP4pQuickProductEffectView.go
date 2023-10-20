package scbp

import (
	"sync"
)

// TopP4pQuickProductEffectView 结构体
type TopP4pQuickProductEffectView struct {
	// 产品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 平均点击花费
	Cpc string `json:"cpc,omitempty" xml:"cpc,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 消耗金额，单位元
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 点击量
	ClickCnt string `json:"click_cnt,omitempty" xml:"click_cnt,omitempty"`
	// 曝光量
	ImpressionCnt string `json:"impression_cnt,omitempty" xml:"impression_cnt,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolTopP4pQuickProductEffectView = sync.Pool{
	New: func() any {
		return new(TopP4pQuickProductEffectView)
	},
}

// GetTopP4pQuickProductEffectView() 从对象池中获取TopP4pQuickProductEffectView
func GetTopP4pQuickProductEffectView() *TopP4pQuickProductEffectView {
	return poolTopP4pQuickProductEffectView.Get().(*TopP4pQuickProductEffectView)
}

// ReleaseTopP4pQuickProductEffectView 释放TopP4pQuickProductEffectView
func ReleaseTopP4pQuickProductEffectView(v *TopP4pQuickProductEffectView) {
	v.ProductName = ""
	v.Cpc = ""
	v.Ctr = ""
	v.Cost = ""
	v.ClickCnt = ""
	v.ImpressionCnt = ""
	v.ProductId = 0
	poolTopP4pQuickProductEffectView.Put(v)
}
