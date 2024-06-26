package logistic

import (
	"sync"
)

// TaobaoWlbImportThreeplOfflineConsignTopResult 结构体
type TaobaoWlbImportThreeplOfflineConsignTopResult struct {
	// 发货完成后的物流单号
	LgOrderCode string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
	// 错误信息
	SubErrorMsg string `json:"sub_error_msg,omitempty" xml:"sub_error_msg,omitempty"`
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误code
	SubErrorCode string `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoWlbImportThreeplOfflineConsignTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoWlbImportThreeplOfflineConsignTopResult)
	},
}

// GetTaobaoWlbImportThreeplOfflineConsignTopResult() 从对象池中获取TaobaoWlbImportThreeplOfflineConsignTopResult
func GetTaobaoWlbImportThreeplOfflineConsignTopResult() *TaobaoWlbImportThreeplOfflineConsignTopResult {
	return poolTaobaoWlbImportThreeplOfflineConsignTopResult.Get().(*TaobaoWlbImportThreeplOfflineConsignTopResult)
}

// ReleaseTaobaoWlbImportThreeplOfflineConsignTopResult 释放TaobaoWlbImportThreeplOfflineConsignTopResult
func ReleaseTaobaoWlbImportThreeplOfflineConsignTopResult(v *TaobaoWlbImportThreeplOfflineConsignTopResult) {
	v.LgOrderCode = ""
	v.SubErrorMsg = ""
	v.ErrorCode = ""
	v.SubErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTaobaoWlbImportThreeplOfflineConsignTopResult.Put(v)
}
