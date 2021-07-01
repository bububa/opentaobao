package nrt

// ResultDo 
type ResultDo struct {
    // succ
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
    // data
    Data   *InvSingleItemSyncDto `json:"data,omitempty" xml:"data,omitempty"`
    // 错误码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 错误信息
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // 系统自动生成
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 系统自动生成
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
