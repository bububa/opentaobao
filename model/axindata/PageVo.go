package axindata

// PageVo 结构体
type PageVo struct {
	// 标准酒店列表
	DataList []StdHotelVo `json:"data_list,omitempty" xml:"data_list>std_hotel_vo,omitempty"`
	// 记录总条数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
