package wdk

// BmPageResult 
/* model for simplify = false
type BmPageResult struct {

    // 当前页码
    
    Current   int64 `json:"current,omitempty"`
    

    // 总条数
    
    Total   int64 `json:"total,omitempty"`
    

    // 对应data
    
    Data  struct {
        PaiyangStatDataResult  []PaiyangStatDataResult `json:"paiyang_stat_data_result,omitempty"`
    } `json:"data,omitempty"`
    

    // 成功
    
    Success   bool `json:"success,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

}
*/

// BmPageResult 
type BmPageResult struct {

    // 当前页码
    Current   int64 `json:"current,omitempty"`

    // 总条数
    Total   int64 `json:"total,omitempty"`

    // 对应data
    Data   []PaiyangStatDataResult `json:"data,omitempty"`

    // 成功
    Success   bool `json:"success,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

}
