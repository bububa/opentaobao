package nrt

// ResultDO 
type ResultDO struct {
    // 是否成功
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
    // 数据
    Data   *DeviceDTO `json:"data,omitempty" xml:"data,omitempty"`
    // 错误码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 测试
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // 错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 调用是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
