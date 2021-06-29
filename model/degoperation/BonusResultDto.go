package degoperation

// BonusResultDTO 
type BonusResultDTO struct {
    // updateAddress=是否填写了收货地址
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // error
    Error   bool `json:"error,omitempty" xml:"error,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
