package wdk

// AlibabaWdkSkuBarcodeQueryApiResults 结构体
type AlibabaWdkSkuBarcodeQueryApiResults struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 条码结果集合
	Models []BarcodeBo `json:"models,omitempty" xml:"models>barcode_bo,omitempty"`
	// 是否成功，根据该字段判断是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
