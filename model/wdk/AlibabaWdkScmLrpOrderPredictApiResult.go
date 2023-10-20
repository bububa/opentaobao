package wdk

import (
	"sync"
)

// AlibabaWdkScmLrpOrderPredictApiResult 结构体
type AlibabaWdkScmLrpOrderPredictApiResult struct {
	// 单量预测结果列表
	PredictList []OrderPredict `json:"predict_list,omitempty" xml:"predict_list>order_predict,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误消息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkScmLrpOrderPredictApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkScmLrpOrderPredictApiResult)
	},
}

// GetAlibabaWdkScmLrpOrderPredictApiResult() 从对象池中获取AlibabaWdkScmLrpOrderPredictApiResult
func GetAlibabaWdkScmLrpOrderPredictApiResult() *AlibabaWdkScmLrpOrderPredictApiResult {
	return poolAlibabaWdkScmLrpOrderPredictApiResult.Get().(*AlibabaWdkScmLrpOrderPredictApiResult)
}

// ReleaseAlibabaWdkScmLrpOrderPredictApiResult 释放AlibabaWdkScmLrpOrderPredictApiResult
func ReleaseAlibabaWdkScmLrpOrderPredictApiResult(v *AlibabaWdkScmLrpOrderPredictApiResult) {
	v.PredictList = v.PredictList[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkScmLrpOrderPredictApiResult.Put(v)
}
