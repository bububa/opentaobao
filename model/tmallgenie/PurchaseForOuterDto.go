package tmallgenie

// PurchaseForOuterDto 结构体
type PurchaseForOuterDto struct {
	// isv用户信息(userId和phone)
	IsvUserInfoMap string `json:"isv_user_info_map,omitempty" xml:"isv_user_info_map,omitempty"`
	// 圈选对象
	CircleInfo *PurchaseCircleInfoForOuterDto `json:"circle_info,omitempty" xml:"circle_info,omitempty"`
	// isvId,由采销系统分配
	IsvId int64 `json:"isv_id,omitempty" xml:"isv_id,omitempty"`
	// 通过查询返回
	SendPlanId int64 `json:"send_plan_id,omitempty" xml:"send_plan_id,omitempty"`
}
