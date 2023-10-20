package wdkitem

// AlibabaWdkItemFuturepriceQueryResult 结构体
type AlibabaWdkItemFuturepriceQueryResult struct {
	// 商品未来价
	Models []AlibabaWdkItemFuturepriceQueryModel `json:"models,omitempty" xml:"models>alibaba_wdk_item_futureprice_query_model,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 异常信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
