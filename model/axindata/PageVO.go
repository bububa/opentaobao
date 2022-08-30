package axindata

// PageVO 结构体
type PageVO struct {
	// 标准酒店列表
	DataList []StdHotelVO `json:"data_list,omitempty" xml:"data_list>std_hotel_vo,omitempty"`
	// 记录总条数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
