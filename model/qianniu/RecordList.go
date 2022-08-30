package qianniu

// RecordList 结构体
type RecordList struct {
	// 变更时间
	ChangeTime string `json:"change_time,omitempty" xml:"change_time,omitempty"`
	// 登陆域
	Domain string `json:"domain,omitempty" xml:"domain,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// mainAccountId
	MainAccountId int64 `json:"main_account_id,omitempty" xml:"main_account_id,omitempty"`
	// 操作记录，1在线，-1离线，2挂起，3解挂
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 记录类型，PC上下线：8；移动上下线：4；挂起类型：5
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 查询记录的帐号ID
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 变更时间戳
	ChangeTimeTs int64 `json:"change_time_ts,omitempty" xml:"change_time_ts,omitempty"`
}
