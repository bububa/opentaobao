package btrip

// CityCarApplyQueryRq 结构体
type CityCarApplyQueryRq struct {
	// 第三方企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 审批单创建时间小于值
	CreatedEndAt string `json:"created_end_at,omitempty" xml:"created_end_at,omitempty"`
	// 审批单创建时间大于等于值
	CreatedStartAt string `json:"created_start_at,omitempty" xml:"created_start_at,omitempty"`
	// 页码，要求大于等于1，默认1
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 每页数据量，要求大于等于1，默认20
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 三方审批单ID
	ThirdPartApplyId string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// 三方员工ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
