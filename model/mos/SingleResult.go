package mos

// SingleResult 
type SingleResult struct {

    // 成功返回
    
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 系统错误
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 系统错误
    
    ErrMessage   string `json:"err_message,omitempty" xml:"err_message,omitempty"`
    

}
