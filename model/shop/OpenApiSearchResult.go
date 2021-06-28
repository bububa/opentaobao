package shop

// OpenApiSearchResult 
/* model for simplify = false
type OpenApiSearchResult struct {

    // 固定值
    
    BucketId   string `json:"bucket_id,omitempty"`
    

    // 固定值
    
    Context   string `json:"context,omitempty"`
    

    // 个数
    
    Count   int64 `json:"count,omitempty"`
    

    // 扩展值
    
    ExtAttrs   string `json:"ext_attrs,omitempty"`
    

    // 店铺信息
    
    Items  struct {
        OpenApiHit  []OpenApiHit `json:"open_api_hit,omitempty"`
    } `json:"items,omitempty"`
    

    // 关键词
    
    Query   string `json:"query,omitempty"`
    

    // 请求ID
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 结果ID
    
    ResultCode   int64 `json:"result_code,omitempty"`
    

    // 结果
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 固定值
    
    SearchGroup   string `json:"search_group,omitempty"`
    

    // 搜索ID
    
    SearchId   string `json:"search_id,omitempty"`
    

    // 固定值
    
    SearchParams   string `json:"search_params,omitempty"`
    

    // sessionID
    
    SessionId   string `json:"session_id,omitempty"`
    

    // 店铺个数
    
    Total   int64 `json:"total,omitempty"`
    

}
*/

// OpenApiSearchResult 
type OpenApiSearchResult struct {

    // 固定值
    BucketId   string `json:"bucket_id,omitempty"`

    // 固定值
    Context   string `json:"context,omitempty"`

    // 个数
    Count   int64 `json:"count,omitempty"`

    // 扩展值
    ExtAttrs   string `json:"ext_attrs,omitempty"`

    // 店铺信息
    Items   []OpenApiHit `json:"items,omitempty"`

    // 关键词
    Query   string `json:"query,omitempty"`

    // 请求ID
    RequestId   string `json:"request_id,omitempty"`

    // 结果ID
    ResultCode   int64 `json:"result_code,omitempty"`

    // 结果
    ResultMsg   string `json:"result_msg,omitempty"`

    // 固定值
    SearchGroup   string `json:"search_group,omitempty"`

    // 搜索ID
    SearchId   string `json:"search_id,omitempty"`

    // 固定值
    SearchParams   string `json:"search_params,omitempty"`

    // sessionID
    SessionId   string `json:"session_id,omitempty"`

    // 店铺个数
    Total   int64 `json:"total,omitempty"`

}
