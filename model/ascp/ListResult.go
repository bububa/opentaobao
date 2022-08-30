package ascp

// ListResult 结构体
type ListResult struct {
	// 库存返回模型
	Result []HiErpQueryInventoryResp `json:"result,omitempty" xml:"result>hi_erp_query_inventory_resp,omitempty"`
	// 错误原因，通常用于success字段值为false时
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误编码，通常用于success为false时的页面错误类型判定
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 通常用于success为true时的页面提示
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 结果是否正确   true：成功  false：失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
