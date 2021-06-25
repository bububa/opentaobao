package hotel

// TopHotelSearchListVo 
type TopHotelSearchListVo struct {

    // 酒店List
    HotelList   []HotelListInfo `json:"hotel_list,omitempty"`

    // 分页参数。下一页开始位置
    Offset   int64 `json:"offset,omitempty"`

    // 总酒店数
    Total   int64 `json:"total,omitempty"`

    // searchId，debug用
    SearchId   string `json:"search_id,omitempty"`

    // totalPages
    TotalPages   int64 `json:"total_pages,omitempty"`

    // 统计信息
    StatInfo   *StatInfo `json:"stat_info,omitempty"`

}
