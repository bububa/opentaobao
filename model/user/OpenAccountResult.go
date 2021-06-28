package user

// OpenAccountResult 
/* model for simplify = false
type OpenAccountResult struct {

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

    // 是否成功
    
    Successful   bool `json:"successful,omitempty"`
    

    // Open Account信息
    
    Data  *struct {
        OpenAccount  *OpenAccount `json:"open_account,omitempty"`
    } `json:"data,omitempty"`
    

    // 错误码
    
    Code   int64 `json:"code,omitempty"`
    

}
*/

// OpenAccountResult 
type OpenAccountResult struct {

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 是否成功
    Successful   bool `json:"successful,omitempty"`

    // Open Account信息
    Data   *OpenAccount `json:"data,omitempty"`

    // 错误码
    Code   int64 `json:"code,omitempty"`

}
