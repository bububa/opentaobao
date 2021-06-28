package alicom

// TopResultDto 
/* model for simplify = false
type TopResultDto struct {

    // 错误信息
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // data
    
    Data  *struct {
        ProductActivityInfoResponseDto  *ProductActivityInfoResponseDto `json:"product_activity_info_response_dto,omitempty"`
    } `json:"data,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TopResultDto 
type TopResultDto struct {

    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty"`

    // data
    Data   *ProductActivityInfoResponseDto `json:"data,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
