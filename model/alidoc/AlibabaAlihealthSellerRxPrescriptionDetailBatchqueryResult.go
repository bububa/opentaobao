package alidoc

import (
	"sync"
)

// AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult 结构体
type AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult struct {
	// 订单处方详情列表
	DataList *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel `json:"data_list,omitempty" xml:"data_list,omitempty"`
}

var poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult)
	},
}

// GetAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult() 从对象池中获取AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult
func GetAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult() *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult {
	return poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult.Get().(*AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult)
}

// ReleaseAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult 释放AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult
func ReleaseAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult(v *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult) {
	v.DataList = nil
	poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult.Put(v)
}
