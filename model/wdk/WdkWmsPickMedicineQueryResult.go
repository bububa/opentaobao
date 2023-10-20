package wdk

import (
	"sync"
)

// WdkWmsPickMedicineQueryResult 结构体
type WdkWmsPickMedicineQueryResult struct {
	// 拣货单维度药品信息list
	MedicineItems []MedicineItemDo `json:"medicine_items,omitempty" xml:"medicine_items>medicine_item_do,omitempty"`
	// 履约单维度药品明细
	SourceOrderMedicineDTOS []SourceOrderMedicineDto `json:"source_order_medicine_d_t_o_s,omitempty" xml:"source_order_medicine_d_t_o_s>source_order_medicine_dto,omitempty"`
	// 系统自动生成
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolWdkWmsPickMedicineQueryResult = sync.Pool{
	New: func() any {
		return new(WdkWmsPickMedicineQueryResult)
	},
}

// GetWdkWmsPickMedicineQueryResult() 从对象池中获取WdkWmsPickMedicineQueryResult
func GetWdkWmsPickMedicineQueryResult() *WdkWmsPickMedicineQueryResult {
	return poolWdkWmsPickMedicineQueryResult.Get().(*WdkWmsPickMedicineQueryResult)
}

// ReleaseWdkWmsPickMedicineQueryResult 释放WdkWmsPickMedicineQueryResult
func ReleaseWdkWmsPickMedicineQueryResult(v *WdkWmsPickMedicineQueryResult) {
	v.MedicineItems = v.MedicineItems[:0]
	v.SourceOrderMedicineDTOS = v.SourceOrderMedicineDTOS[:0]
	v.ErrorMsg = ""
	v.Success = false
	poolWdkWmsPickMedicineQueryResult.Put(v)
}
