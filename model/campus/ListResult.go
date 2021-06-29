package campus

// ListResult 
type ListResult struct {
    // content
    Contents   []string `json:"contents,omitempty" xml:"contents>string,omitempty"`
    // 总记录数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否调用成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // requestId
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // errorLevel
    ErrorLevel   string `json:"error_level,omitempty" xml:"error_level,omitempty"`
    // 设备历史数据内容
    DeviceDataList   []DeviceHistoryBatchApiDto `json:"device_data_list,omitempty" xml:"device_data_list>device_history_batch_api_dto,omitempty"`
    // 额外报错信息
    ErrorExtInfo   string `json:"error_ext_info,omitempty" xml:"error_ext_info,omitempty"`
    // 规则列表数据
    ContentList   []string `json:"content_list,omitempty" xml:"content_list>string,omitempty"`
}
