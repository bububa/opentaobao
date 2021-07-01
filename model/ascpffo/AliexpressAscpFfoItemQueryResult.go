package ascpffo

// AliexpressAscpFfoItemQueryResult 结构体
type AliexpressAscpFfoItemQueryResult struct {
	// dto
	DataList []AliexpressAscpFfoItemQueryData `json:"data_list,omitempty" xml:"data_list>aliexpress_ascp_ffo_item_query_data,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
