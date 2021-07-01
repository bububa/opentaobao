package ascpffo

// PageQueryResult 
type PageQueryResult struct {
    // dto list
    DataList   []ErpFulfillmentForwardDto `json:"data_list,omitempty" xml:"data_list>erp_fulfillment_forward_dto,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 分页页码
    PageIndex   int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
    // 分页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 数量
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
