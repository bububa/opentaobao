package util

// ItemPickPagingResult 
type ItemPickPagingResult struct {
    // 返回数据集合
    Results   []CountryDto `json:"results,omitempty" xml:"results>country_dto,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 页码
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 每页数据条数
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 总记录条数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
