package opentrade

// GroupDetailResponse 结构体
type GroupDetailResponse struct {
	// 已参团人数
	Joined int64 `json:"joined,omitempty" xml:"joined,omitempty"`
	// 成团目标数
	Goal int64 `json:"goal,omitempty" xml:"goal,omitempty"`
	// 组团类型，0 拼团 1 团购
	GroupType int64 `json:"group_type,omitempty" xml:"group_type,omitempty"`
	// 团id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 团长信息
	Leader *RopGroupMemberInfo `json:"leader,omitempty" xml:"leader,omitempty"`
	// 团员信息
	Members []RopGroupMemberInfo `json:"members,omitempty" xml:"members>rop_group_member_info,omitempty"`
	// 是否成团
	Done bool `json:"done,omitempty" xml:"done,omitempty"`
	// 成团有效截止日期
	Expiration string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// 下单付款人数
	Payed int64 `json:"payed,omitempty" xml:"payed,omitempty"`
	// 组团活动id
	GroupActivityId int64 `json:"group_activity_id,omitempty" xml:"group_activity_id,omitempty"`
	// 1-未付款，2-已付款，4-已退款，6-交易成功，8-交易关闭
	MemberPayStatus int64 `json:"member_pay_status,omitempty" xml:"member_pay_status,omitempty"`
}
