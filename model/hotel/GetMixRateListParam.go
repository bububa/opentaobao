package hotel

// GetMixRateListParam 
/* model for simplify = false
type GetMixRateListParam struct {

    // 业务类型
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 用于嵌入页截断显示
    
    Limit   int64 `json:"limit,omitempty"`
    

    // 是否加载顶踩数据，为0不加载，为1加载
    
    LoadAttitude   int64 `json:"load_attitude,omitempty"`
    

    // 是否加载回复数据，为0不加载，为1加载
    
    LoadReply   int64 `json:"load_reply,omitempty"`
    

    // 页面号，第几页
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 页面包含的记录数
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 酒店评论类型筛选
    
    TabFilter   string `json:"tab_filter,omitempty"`
    

}
*/

// GetMixRateListParam 
type GetMixRateListParam struct {

    // 业务类型
    ItemId   int64 `json:"item_id,omitempty"`

    // 用于嵌入页截断显示
    Limit   int64 `json:"limit,omitempty"`

    // 是否加载顶踩数据，为0不加载，为1加载
    LoadAttitude   int64 `json:"load_attitude,omitempty"`

    // 是否加载回复数据，为0不加载，为1加载
    LoadReply   int64 `json:"load_reply,omitempty"`

    // 页面号，第几页
    PageNo   int64 `json:"page_no,omitempty"`

    // 页面包含的记录数
    PageSize   int64 `json:"page_size,omitempty"`

    // 酒店评论类型筛选
    TabFilter   string `json:"tab_filter,omitempty"`

}
