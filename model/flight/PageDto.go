package flight

// PageDto 
type PageDto struct {

    // 123
    Total   int64 `json:"total,omitempty"`

    // 结果集
    DataList   []T `json:"data_list,omitempty"`

    // 成功标识
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
