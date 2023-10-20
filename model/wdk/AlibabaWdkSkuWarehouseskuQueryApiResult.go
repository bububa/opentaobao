package wdk

// AlibabawdkskuwarehouseskuqueryApiResult 结构体
type AlibabawdkskuwarehouseskuqueryApiResult struct {
	// 数据集合
	Models []WarehouseSkuDo `json:"models,omitempty" xml:"models>warehouse_sku_do,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误内容
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
