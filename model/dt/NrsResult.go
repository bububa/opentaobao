package dt

// NrsResult 
type NrsResult struct {

    // 返回数据
    Data   *RecongnizeItemInfo `json:"data,omitempty"`

    // 接口调用标志
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

}
