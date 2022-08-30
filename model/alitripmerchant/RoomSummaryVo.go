package alitripmerchant

// RoomSummaryVo 结构体
type RoomSummaryVo struct {
	// 房型名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 房型id
	SrId int64 `json:"sr_id,omitempty" xml:"sr_id,omitempty"`
}
