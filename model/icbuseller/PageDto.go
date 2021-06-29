package icbuseller

// PageDTO 
type PageDTO struct {
    // 总数据量
    TotalItem   int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
    // 总页数
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
    // 每页显示数量
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 当前页码
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
