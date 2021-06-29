package store

// TopBatchResultDO 
type TopBatchResultDO struct {
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否失败
    Failure   bool `json:"failure,omitempty" xml:"failure,omitempty"`
    // 其它信息
    Other   *Other `json:"other,omitempty" xml:"other,omitempty"`
    // 失败的门店id
    ResultList   []int64 `json:"result_list,omitempty" xml:"result_list>int64,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 总条目数
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
