package traderate

// TopGetMixRateListParam 
/* model for simplify = false
type TopGetMixRateListParam struct {

    // 类型
    
    BizType   int64 `json:"biz_type,omitempty"`
    

    // 酒店或商品ID
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // pageNo
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // pageSize
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 筛选条件JSON
    
    TabFilter   string `json:"tab_filter,omitempty"`
    

}
*/

// TopGetMixRateListParam 
type TopGetMixRateListParam struct {

    // 类型
    BizType   int64 `json:"biz_type,omitempty"`

    // 酒店或商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // pageNo
    PageNo   int64 `json:"page_no,omitempty"`

    // pageSize
    PageSize   int64 `json:"page_size,omitempty"`

    // 筛选条件JSON
    TabFilter   string `json:"tab_filter,omitempty"`

}
