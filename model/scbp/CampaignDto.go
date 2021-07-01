package scbp

// CampaignDto 结构体
type CampaignDto struct {
	// 计划标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 计划开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 计划结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 用户上下线
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 结算上下线
	SettleStatus int64 `json:"settle_status,omitempty" xml:"settle_status,omitempty"`
	// 结算上下线原因
	SettleReason string `json:"settle_reason,omitempty" xml:"settle_reason,omitempty"`
	// settleTime
	SettleTime string `json:"settle_time,omitempty" xml:"settle_time,omitempty"`
	// 计划类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 计划模型
	CampaignModel int64 `json:"campaign_model,omitempty" xml:"campaign_model,omitempty"`
	// 置顶时间
	TopTime string `json:"top_time,omitempty" xml:"top_time,omitempty"`
	// 结算版本
	SettleVersion int64 `json:"settle_version,omitempty" xml:"settle_version,omitempty"`
	// 场景ID
	SceneId int64 `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	// Campaign扩展属性
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
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
	// 计划创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 计划修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}
