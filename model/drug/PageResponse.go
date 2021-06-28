package drug

// PageResponse 
/* model for simplify = false
type PageResponse struct {

    // 当前页码
    
    CurrentPage   int64 `json:"current_page,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 总条数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 结果列表
    
    Spus  struct {
        TopAlihealthSpuQuery  []TopAlihealthSpuQuery `json:"top_alihealth_spu_query,omitempty"`
    } `json:"spus,omitempty"`
    

}
*/

// PageResponse 
type PageResponse struct {

    // 当前页码
    CurrentPage   int64 `json:"current_page,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 总条数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 结果列表
    Spus   []TopAlihealthSpuQuery `json:"spus,omitempty"`

}
