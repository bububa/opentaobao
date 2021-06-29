package lsttrade

// AlibabaLstTradeSellerOrderListQueryResult 
type AlibabaLstTradeSellerOrderListQueryResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 失败信息
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 系统自动生成
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 数量
    Size   int64 `json:"size,omitempty" xml:"size,omitempty"`
    // 当前页
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    // 信息
    ContentList   []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
}
