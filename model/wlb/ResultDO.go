package wlb

// ResultDO 
type ResultDO struct {

    // 网络延时
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 01
    ErrorCode   string `json:"error_code,omitempty"`

    // 成功、失败
    Success   bool `json:"success,omitempty"`

}
