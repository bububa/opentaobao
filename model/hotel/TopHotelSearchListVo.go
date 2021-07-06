package hotel

// TopHotelSearchListVo 结构体
type TopHotelSearchListVo struct {
	// 酒店List
	HotelList []HotelListInfo `json:"hotel_list,omitempty" xml:"hotel_list>hotel_list_info,omitempty"`
	// searchId，debug用
	SearchId string `json:"search_id,omitempty" xml:"search_id,omitempty"`
	// 分页参数。下一页开始位置
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 总酒店数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// totalPages
	TotalPages int64 `json:"total_pages,omitempty" xml:"total_pages,omitempty"`
	// 统计信息
	StatInfo *StatInfo `json:"stat_info,omitempty" xml:"stat_info,omitempty"`
}
