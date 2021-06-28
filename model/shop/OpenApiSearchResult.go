package shop

// OpenApiSearchResult 
type OpenApiSearchResult struct {

    // 固定值
    
    BucketId   string `json:"bucket_id,omitempty" xml:"bucket_id,omitempty"`
    

    // 固定值
    
    Context   string `json:"context,omitempty" xml:"context,omitempty"`
    

    // 个数
    
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    

    // 扩展值
    
    ExtAttrs   string `json:"ext_attrs,omitempty" xml:"ext_attrs,omitempty"`
    

    // 店铺信息
    
    Items   []OpenApiHit `json:"items,omitempty" xml:"items,omitempty"`
    

    // 关键词
    
    Query   string `json:"query,omitempty" xml:"query,omitempty"`
    

    // 请求ID
    
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    

    // 结果ID
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
    

    // 结果
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    

    // 固定值
    
    SearchGroup   string `json:"search_group,omitempty" xml:"search_group,omitempty"`
    

    // 搜索ID
    
    SearchId   string `json:"search_id,omitempty" xml:"search_id,omitempty"`
    

    // 固定值
    
    SearchParams   string `json:"search_params,omitempty" xml:"search_params,omitempty"`
    

    // sessionID
    
    SessionId   string `json:"session_id,omitempty" xml:"session_id,omitempty"`
    

    // 店铺个数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

}
