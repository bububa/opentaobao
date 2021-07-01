package wdk

// AlibabaWdkScmLrpOrderPredictApiResult 结构体
type AlibabaWdkScmLrpOrderPredictApiResult struct {
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误消息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 单量预测结果列表
	PredictList []OrderPredict `json:"predict_list,omitempty" xml:"predict_list>order_predict,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
