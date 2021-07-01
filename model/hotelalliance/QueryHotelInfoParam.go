package hotelalliance

// QueryHotelInfoParam 结构体
type QueryHotelInfoParam struct {
	// 飞猪卖家酒店id
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 单体联盟飞猪卖家酒店id
	Ahid int64 `json:"ahid,omitempty" xml:"ahid,omitempty"`
}
