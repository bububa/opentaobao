package lstwarehouse

// AlibabaLstBranddatashareStockdataQueryResult 
type AlibabaLstBranddatashareStockdataQueryResult struct {
    // 总记录数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 每页记录数
    Size   int64 `json:"size,omitempty" xml:"size,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误描述
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 页码
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    // 返回内容
    ContentList   []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
}
