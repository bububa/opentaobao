package alitripmerchant

// PageableResponse 
type PageableResponse struct {
    // 当前页的数量
    Size   int64 `json:"size,omitempty" xml:"size,omitempty"`
    // 是否有下一页
    HasNextPage   bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
    // 当前页
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 每页的数量
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 错误编码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 总页数
    TotalPageNum   int64 `json:"total_page_num,omitempty" xml:"total_page_num,omitempty"`
    // 总数量
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 返回类型
    Contents   []HotelListSearchDTO `json:"contents,omitempty" xml:"contents>hotel_list_search_dto,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
