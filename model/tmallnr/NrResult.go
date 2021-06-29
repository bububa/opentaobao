package tmallnr

// NrResult 
type NrResult struct {
    // 取消是否成功
    ResultData   bool `json:"result_data,omitempty" xml:"result_data,omitempty"`
    // errorMessage
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // errorCode
    ErrorCode2   string `json:"error_code2,omitempty" xml:"error_code2,omitempty"`
    // isSuccess
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // isSuccess
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 是否成功
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
    // 系统自动生成
    ResultDatas   []ResultData `json:"result_datas,omitempty" xml:"result_datas>result_data,omitempty"`
}
