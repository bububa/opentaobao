package drugtrace

// Header 
type Header struct {

    // 消息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 标记
    
    SuccessFlag   string `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
    

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

}
