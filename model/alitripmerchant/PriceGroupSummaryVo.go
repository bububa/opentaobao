package alitripmerchant

// PriceGroupSummaryVo 结构体
type PriceGroupSummaryVo struct {
	// 房型信息
	Rooms []RoomSummaryVo `json:"rooms,omitempty" xml:"rooms>room_summary_vo,omitempty"`
	// 组名
	GroupName string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// 最低价格
	LowestPrice int64 `json:"lowest_price,omitempty" xml:"lowest_price,omitempty"`
	// 房间总数
	RoomTotal int64 `json:"room_total,omitempty" xml:"room_total,omitempty"`
}
