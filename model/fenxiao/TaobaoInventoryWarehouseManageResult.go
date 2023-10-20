package fenxiao

import (
	"sync"
)

// TaobaoInventoryWarehouseManageResult 结构体
type TaobaoInventoryWarehouseManageResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoInventoryWarehouseManageResult = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryWarehouseManageResult)
	},
}

// GetTaobaoInventoryWarehouseManageResult() 从对象池中获取TaobaoInventoryWarehouseManageResult
func GetTaobaoInventoryWarehouseManageResult() *TaobaoInventoryWarehouseManageResult {
	return poolTaobaoInventoryWarehouseManageResult.Get().(*TaobaoInventoryWarehouseManageResult)
}

// ReleaseTaobaoInventoryWarehouseManageResult 释放TaobaoInventoryWarehouseManageResult
func ReleaseTaobaoInventoryWarehouseManageResult(v *TaobaoInventoryWarehouseManageResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = false
	v.Success = false
	poolTaobaoInventoryWarehouseManageResult.Put(v)
}
