package alitripmerchant

// HotelDetailsParam 结构体
type HotelDetailsParam struct {
	// 分组过滤请求
	GroupParam []HotelDetailsGroupParam `json:"group_param,omitempty" xml:"group_param>hotel_details_group_param,omitempty"`
	// 离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 入店时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 外部酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 会员等级
	MemberLevel string `json:"member_level,omitempty" xml:"member_level,omitempty"`
	// 儿童年龄数组
	ChildrenAges string `json:"children_ages,omitempty" xml:"children_ages,omitempty"`
	// 快筛选项
	RoomType string `json:"room_type,omitempty" xml:"room_type,omitempty"`
	// 用户登录信息
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 代金券id
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 成人数量
	AdultNum int64 `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
	// 连住几晚
	FewNights int64 `json:"few_nights,omitempty" xml:"few_nights,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
