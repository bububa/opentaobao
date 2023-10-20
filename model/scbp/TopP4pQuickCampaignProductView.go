package scbp

import (
	"sync"
)

// TopP4pQuickCampaignProductView 结构体
type TopP4pQuickCampaignProductView struct {
	// 商品名
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 商品最近7天效果数据
	Effect7d *Effect7d `json:"effect7d,omitempty" xml:"effect7d,omitempty"`
	// 商品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 产品状态（0暂停，1推广中，-2商品下架）
	DisplayStatus int64 `json:"display_status,omitempty" xml:"display_status,omitempty"`
}

var poolTopP4pQuickCampaignProductView = sync.Pool{
	New: func() any {
		return new(TopP4pQuickCampaignProductView)
	},
}

// GetTopP4pQuickCampaignProductView() 从对象池中获取TopP4pQuickCampaignProductView
func GetTopP4pQuickCampaignProductView() *TopP4pQuickCampaignProductView {
	return poolTopP4pQuickCampaignProductView.Get().(*TopP4pQuickCampaignProductView)
}

// ReleaseTopP4pQuickCampaignProductView 释放TopP4pQuickCampaignProductView
func ReleaseTopP4pQuickCampaignProductView(v *TopP4pQuickCampaignProductView) {
	v.ProductName = ""
	v.Effect7d = nil
	v.ProductId = 0
	v.DisplayStatus = 0
	poolTopP4pQuickCampaignProductView.Put(v)
}
