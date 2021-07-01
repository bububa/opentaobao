package btrip

// OpenApproverInfo 结构体
type OpenApproverInfo struct {
	// 操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 审批意见
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// 审批人状态：0审批中 1已同意 2已拒绝 3已转交，4已取消 5已终止 6发起审批 7评论
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 审批人名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 审批人id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 审批人顺序
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 审批人状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
}
