package mos

// AlibabaMjOcSyncpayinfoResultDo 
type AlibabaMjOcSyncpayinfoResultDo struct {

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 错误码
    ErrCode   int64 `json:"err_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // data
    Date   string `json:"date,omitempty"`

}