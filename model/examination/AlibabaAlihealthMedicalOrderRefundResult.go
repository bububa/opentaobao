package examination

import (
	"sync"
)

// AlibabaAlihealthMedicalOrderRefundResult 结构体
type AlibabaAlihealthMedicalOrderRefundResult struct {
	// SUCCESS:成功; FAIL:失败; UNKNOWN:未知;
	ResultStatus string `json:"result_status,omitempty" xml:"result_status,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回数据
	Data *OrderRefundVo `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihealthMedicalOrderRefundResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalOrderRefundResult)
	},
}

// GetAlibabaAlihealthMedicalOrderRefundResult() 从对象池中获取AlibabaAlihealthMedicalOrderRefundResult
func GetAlibabaAlihealthMedicalOrderRefundResult() *AlibabaAlihealthMedicalOrderRefundResult {
	return poolAlibabaAlihealthMedicalOrderRefundResult.Get().(*AlibabaAlihealthMedicalOrderRefundResult)
}

// ReleaseAlibabaAlihealthMedicalOrderRefundResult 释放AlibabaAlihealthMedicalOrderRefundResult
func ReleaseAlibabaAlihealthMedicalOrderRefundResult(v *AlibabaAlihealthMedicalOrderRefundResult) {
	v.ResultStatus = ""
	v.ResultCode = ""
	v.ResultMsg = ""
	v.Data = nil
	poolAlibabaAlihealthMedicalOrderRefundResult.Put(v)
}
