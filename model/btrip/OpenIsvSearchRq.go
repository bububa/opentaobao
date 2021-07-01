package btrip

// OpenIsvSearchRq 结构体
type OpenIsvSearchRq struct {
	// 阿里商旅审批单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 阿里商旅审批单展示id
	ApplyShowId string `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	// 第三方企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 外部审批单id
	ThirdpartApplyId string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// 1、老版本2、isv对外版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 申请单提交类型 1代提交 2本人提交 注意：当申请单为代提交时，申请单提交人自己无法为自己下单
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// false:未报销的申请单
	AllApply bool `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	// 申请人所在部门id
	DepartId string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 更新时间大于等于此时间的审批单
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// true:商旅申请单
	OnlyShangLvApply bool `json:"only_shang_lv_apply,omitempty" xml:"only_shang_lv_apply,omitempty"`
	// 页数，从1开始
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页返回数量，默认10，最多50
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 申请人Id（第三方用户id）
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
