package btrip

// CityCarApplyApproveRq 结构体
type CityCarApplyApproveRq struct {
	// 第三方企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 审批时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 审批备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 审批结果：1-同意，2-拒绝
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 第三方审批单ID
	ThirdPartApplyId string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// 审批的第三方员工ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
