package store

// TopResultDo 
type TopResultDo struct {
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否失败
    Failure   bool `json:"failure,omitempty" xml:"failure,omitempty"`
    // 返回结果：true成功；false失败
    Result   *FullStoreTopDto `json:"result,omitempty" xml:"result,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 个数
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
