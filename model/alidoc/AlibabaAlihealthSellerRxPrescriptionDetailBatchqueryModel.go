package alidoc

import (
	"sync"
)

// AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel 结构体
type AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel struct {
	// 未匹配到的订单请求列表
	FailedList []FailedList `json:"failed_list,omitempty" xml:"failed_list>failed_list,omitempty"`
	// 匹配到的订单处方详情列表
	PrescriptionDetailList []PrescriptionDetailTopDto `json:"prescription_detail_list,omitempty" xml:"prescription_detail_list>prescription_detail_top_dto,omitempty"`
}

var poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel)
	},
}

// GetAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel() 从对象池中获取AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel
func GetAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel() *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel {
	return poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel.Get().(*AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel)
}

// ReleaseAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel 释放AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel
func ReleaseAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel(v *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel) {
	v.FailedList = v.FailedList[:0]
	v.PrescriptionDetailList = v.PrescriptionDetailList[:0]
	poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryModel.Put(v)
}
