package fenxiao

// PaginationResult 
/* model for simplify = false
type PaginationResult struct {

    // 总条数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 仓库信息数组
    
    Data   string `json:"data,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 页码
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// PaginationResult 
type PaginationResult struct {

    // 总条数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 仓库信息数组
    Data   string `json:"data,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 页码
    PageNo   int64 `json:"page_no,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
