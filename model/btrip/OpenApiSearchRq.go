package btrip

// OpenApiSearchRq 结构体
type OpenApiSearchRq struct {
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 用户所在部门id
	DepartId string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 第三方申请单ID
	ThirdpartApplyId string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// 更新结束时间
	UpdateEndTime string `json:"update_end_time,omitempty" xml:"update_end_time,omitempty"`
	// 更新开始时间
	UpdateStartTime string `json:"update_start_time,omitempty" xml:"update_start_time,omitempty"`
	// 第三方用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 商旅申请单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 页数，从1开始
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数据量，默认10，最高50
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 1、老版本2、isv对外版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// false：仅搜索未报销的申请单
	AllApply bool `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
}
