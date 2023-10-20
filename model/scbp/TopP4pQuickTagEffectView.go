package scbp

import (
	"sync"
)

// TopP4pQuickTagEffectView 结构体
type TopP4pQuickTagEffectView struct {
	// 标签id(潜在访问偏好和潜在采购意向返回的是类目id，店铺老客和优选人群返回的是字符串)
	TagId string `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
	// 标签名字
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 平均点击出价
	Cpc string `json:"cpc,omitempty" xml:"cpc,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 消耗金额，单位元
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 点击数
	ClickCnt string `json:"click_cnt,omitempty" xml:"click_cnt,omitempty"`
	// 曝光量
	ImpressionCnt string `json:"impression_cnt,omitempty" xml:"impression_cnt,omitempty"`
}

var poolTopP4pQuickTagEffectView = sync.Pool{
	New: func() any {
		return new(TopP4pQuickTagEffectView)
	},
}

// GetTopP4pQuickTagEffectView() 从对象池中获取TopP4pQuickTagEffectView
func GetTopP4pQuickTagEffectView() *TopP4pQuickTagEffectView {
	return poolTopP4pQuickTagEffectView.Get().(*TopP4pQuickTagEffectView)
}

// ReleaseTopP4pQuickTagEffectView 释放TopP4pQuickTagEffectView
func ReleaseTopP4pQuickTagEffectView(v *TopP4pQuickTagEffectView) {
	v.TagId = ""
	v.TagName = ""
	v.Cpc = ""
	v.Ctr = ""
	v.Cost = ""
	v.ClickCnt = ""
	v.ImpressionCnt = ""
	poolTopP4pQuickTagEffectView.Put(v)
}
