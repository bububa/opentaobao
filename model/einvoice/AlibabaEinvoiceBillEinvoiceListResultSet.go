package einvoice

// AlibabaEinvoiceBillEinvoiceListResultSet 结构体
type AlibabaEinvoiceBillEinvoiceListResultSet struct {
	// 返回结果具体信息
	ResultList []ResultList `json:"result_list,omitempty" xml:"result_list>result_list,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
