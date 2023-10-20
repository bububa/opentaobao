package scbp

import (
	"sync"
)

// CampaignDto 结构体
type CampaignDto struct {
	// 计划标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 计划开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 计划结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 结算上下线原因
	SettleReason string `json:"settle_reason,omitempty" xml:"settle_reason,omitempty"`
	// settleTime
	SettleTime string `json:"settle_time,omitempty" xml:"settle_time,omitempty"`
	// 置顶时间
	TopTime string `json:"top_time,omitempty" xml:"top_time,omitempty"`
	// Campaign扩展属性
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 计划创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 计划修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 计划子类型
	SubType string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
	// 周预算状态 on:开启 off:关闭
	WeekBudgetStatus string `json:"week_budget_status,omitempty" xml:"week_budget_status,omitempty"`
	// 最低价格
	MinPrice string `json:"min_price,omitempty" xml:"min_price,omitempty"`
	// 最大价格
	MaxPrice string `json:"max_price,omitempty" xml:"max_price,omitempty"`
	// 预算
	Budget string `json:"budget,omitempty" xml:"budget,omitempty"`
	// 用户上下线
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 结算上下线
	SettleStatus int64 `json:"settle_status,omitempty" xml:"settle_status,omitempty"`
	// 计划类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 计划模型
	CampaignModel int64 `json:"campaign_model,omitempty" xml:"campaign_model,omitempty"`
	// 结算版本
	SettleVersion int64 `json:"settle_version,omitempty" xml:"settle_version,omitempty"`
	// 场景ID
	SceneId int64 `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	// 产品线id
	ProductLineId int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`
	// 广告用户ID
	MemberId int64 `json:"member_id,omitempty" xml:"member_id,omitempty"`
	// 客户ID
	CustId int64 `json:"cust_id,omitempty" xml:"cust_id,omitempty"`
	// 业务流水号
	BizNumber int64 `json:"biz_number,omitempty" xml:"biz_number,omitempty"`
	// 实体主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 出价模式 value =1 为智能出价 value =2 为手动出价
	BidType int64 `json:"bid_type,omitempty" xml:"bid_type,omitempty"`
}

var poolCampaignDto = sync.Pool{
	New: func() any {
		return new(CampaignDto)
	},
}

// GetCampaignDto() 从对象池中获取CampaignDto
func GetCampaignDto() *CampaignDto {
	return poolCampaignDto.Get().(*CampaignDto)
}

// ReleaseCampaignDto 释放CampaignDto
func ReleaseCampaignDto(v *CampaignDto) {
	v.Title = ""
	v.StartTime = ""
	v.EndTime = ""
	v.SettleReason = ""
	v.SettleTime = ""
	v.TopTime = ""
	v.Properties = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.SubType = ""
	v.WeekBudgetStatus = ""
	v.MinPrice = ""
	v.MaxPrice = ""
	v.Budget = ""
	v.OnlineStatus = 0
	v.SettleStatus = 0
	v.Type = 0
	v.CampaignModel = 0
	v.SettleVersion = 0
	v.SceneId = 0
	v.ProductLineId = 0
	v.MemberId = 0
	v.CustId = 0
	v.BizNumber = 0
	v.Id = 0
	v.BidType = 0
	poolCampaignDto.Put(v)
}
