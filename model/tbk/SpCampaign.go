package tbk

// SpCampaign 结构体
type SpCampaign struct {
	// 定向计划活动ID
	Spcid string `json:"sp_cid,omitempty" xml:"sp_cid,omitempty"`
	// 定向计划名称
	Spname string `json:"sp_name,omitempty" xml:"sp_name,omitempty"`
	// 定向佣金率
	Sprate string `json:"sp_rate,omitempty" xml:"sp_rate,omitempty"`
	// 定向是否锁佣，0=不锁佣 1=锁佣
	Splockstatus string `json:"sp_lock_status,omitempty" xml:"sp_lock_status,omitempty"`
	// 定向计划申请链接
	Spapplylink string `json:"sp_apply_link,omitempty" xml:"sp_apply_link,omitempty"`
	// 定向计划是否可用 1-可用 0-不可用
	Spstatus string `json:"sp_status,omitempty" xml:"sp_status,omitempty"`
}
