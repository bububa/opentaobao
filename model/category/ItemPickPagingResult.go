package category

// ItemPickPagingResult 
/* model for simplify = false
type ItemPickPagingResult struct {

    // result
    
    Results  struct {
        CategoryDto  []CategoryDto `json:"category_dto,omitempty"`
    } `json:"results,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// ItemPickPagingResult 
type ItemPickPagingResult struct {

    // result
    Results   []CategoryDto `json:"results,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
