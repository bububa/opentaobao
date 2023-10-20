package scbp

import (
	"sync"
)

// CampaignOperationDto 结构体
type CampaignOperationDto struct {
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 最小价格
	MinPrice string `json:"min_price,omitempty" xml:"min_price,omitempty"`
	// 最大价格
	MaxPrice string `json:"max_price,omitempty" xml:"max_price,omitempty"`
	// 计划子类型
	SubType string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
	// 预算
	Budget string `json:"budget,omitempty" xml:"budget,omitempty"`
	// 周预算状态 on:开启 off:关闭
	WeekBudgetStatus string `json:"week_budget_status,omitempty" xml:"week_budget_status,omitempty"`
	// 扩展匹配：on：开启 off：关闭
	SyncExtMatchConf string `json:"sync_ext_match_conf,omitempty" xml:"sync_ext_match_conf,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 计划类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 计划状态
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 出价模式  * value =1 为智能出价      * value =2 为手动出价
	BidType int64 `json:"bid_type,omitempty" xml:"bid_type,omitempty"`
	// 计划id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolCampaignOperationDto = sync.Pool{
	New: func() any {
		return new(CampaignOperationDto)
	},
}

// GetCampaignOperationDto() 从对象池中获取CampaignOperationDto
func GetCampaignOperationDto() *CampaignOperationDto {
	return poolCampaignOperationDto.Get().(*CampaignOperationDto)
}

// ReleaseCampaignOperationDto 释放CampaignOperationDto
func ReleaseCampaignOperationDto(v *CampaignOperationDto) {
	v.Title = ""
	v.MinPrice = ""
	v.MaxPrice = ""
	v.SubType = ""
	v.Budget = ""
	v.WeekBudgetStatus = ""
	v.SyncExtMatchConf = ""
	v.StartTime = ""
	v.Type = 0
	v.OnlineStatus = 0
	v.BidType = 0
	v.Id = 0
	poolCampaignOperationDto.Put(v)
}
