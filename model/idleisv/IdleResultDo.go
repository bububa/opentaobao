package idleisv

// IdleResultDO 
type IdleResultDO struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 返回数据
    Data   *IdleItemApiDO `json:"data,omitempty" xml:"data,omitempty"`
    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
}
