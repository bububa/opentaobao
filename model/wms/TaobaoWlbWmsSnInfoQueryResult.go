package wms

// TaobaoWlbWmsSnInfoQueryResult 
type TaobaoWlbWmsSnInfoQueryResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // SN信息列表
    SnInfoList   []Sninfolist `json:"sn_info_list,omitempty"`

    // 总条数
    TotalCount   int64 `json:"total_count,omitempty"`

}
