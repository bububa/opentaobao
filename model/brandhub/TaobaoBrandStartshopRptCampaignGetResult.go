package brandhub

import (
	"sync"
)

// TaobaoBrandStartshopRptCampaignGetResult 结构体
type TaobaoBrandStartshopRptCampaignGetResult struct {
	// 日期
	Thedate string `json:"thedate,omitempty" xml:"thedate,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 花费(元)
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 千次展现成本
	Cpm string `json:"cpm,omitempty" xml:"cpm,omitempty"`
	// 点击单价
	Cpc string `json:"cpc,omitempty" xml:"cpc,omitempty"`
	// 展示成交金额
	Transactiontotal string `json:"transactiontotal,omitempty" xml:"transactiontotal,omitempty"`
	// 展示回报率
	Roi string `json:"roi,omitempty" xml:"roi,omitempty"`
	// 展示转化率
	Cvr string `json:"cvr,omitempty" xml:"cvr,omitempty"`
	// 点击成交金额
	ClickTransactiontotal string `json:"click_transactiontotal,omitempty" xml:"click_transactiontotal,omitempty"`
	// 点击回报率
	ClickRoi string `json:"click_roi,omitempty" xml:"click_roi,omitempty"`
	// 点击转化率
	ClickCvr string `json:"click_cvr,omitempty" xml:"click_cvr,omitempty"`
	// 推广计划名称
	Campaigntitle string `json:"campaigntitle,omitempty" xml:"campaigntitle,omitempty"`
	// 展现量
	Impression int64 `json:"impression,omitempty" xml:"impression,omitempty"`
	// 点击量
	Click int64 `json:"click,omitempty" xml:"click,omitempty"`
	// 宝贝收藏数
	Favitemtotal int64 `json:"favitemtotal,omitempty" xml:"favitemtotal,omitempty"`
	// 展示成交笔数
	Transactionshippingtotal int64 `json:"transactionshippingtotal,omitempty" xml:"transactionshippingtotal,omitempty"`
	// 店铺收藏数
	Favshoptotal int64 `json:"favshoptotal,omitempty" xml:"favshoptotal,omitempty"`
	// 宝贝加购数
	Carttotal int64 `json:"carttotal,omitempty" xml:"carttotal,omitempty"`
	// 点击成交笔数
	ClickTransactionshipping int64 `json:"click_transactionshipping,omitempty" xml:"click_transactionshipping,omitempty"`
	// 点击访客数
	ClickUv int64 `json:"click_uv,omitempty" xml:"click_uv,omitempty"`
	// 触达访客数
	Uv int64 `json:"uv,omitempty" xml:"uv,omitempty"`
	// 触达新访客数
	UvNew int64 `json:"uv_new,omitempty" xml:"uv_new,omitempty"`
	// 推广计划id
	Campaignid int64 `json:"campaignid,omitempty" xml:"campaignid,omitempty"`
}

var poolTaobaoBrandStartshopRptCampaignGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoBrandStartshopRptCampaignGetResult)
	},
}

// GetTaobaoBrandStartshopRptCampaignGetResult() 从对象池中获取TaobaoBrandStartshopRptCampaignGetResult
func GetTaobaoBrandStartshopRptCampaignGetResult() *TaobaoBrandStartshopRptCampaignGetResult {
	return poolTaobaoBrandStartshopRptCampaignGetResult.Get().(*TaobaoBrandStartshopRptCampaignGetResult)
}

// ReleaseTaobaoBrandStartshopRptCampaignGetResult 释放TaobaoBrandStartshopRptCampaignGetResult
func ReleaseTaobaoBrandStartshopRptCampaignGetResult(v *TaobaoBrandStartshopRptCampaignGetResult) {
	v.Thedate = ""
	v.Ctr = ""
	v.Cost = ""
	v.Cpm = ""
	v.Cpc = ""
	v.Transactiontotal = ""
	v.Roi = ""
	v.Cvr = ""
	v.ClickTransactiontotal = ""
	v.ClickRoi = ""
	v.ClickCvr = ""
	v.Campaigntitle = ""
	v.Impression = 0
	v.Click = 0
	v.Favitemtotal = 0
	v.Transactionshippingtotal = 0
	v.Favshoptotal = 0
	v.Carttotal = 0
	v.ClickTransactionshipping = 0
	v.ClickUv = 0
	v.Uv = 0
	v.UvNew = 0
	v.Campaignid = 0
	poolTaobaoBrandStartshopRptCampaignGetResult.Put(v)
}
