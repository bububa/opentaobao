package wdk

import (
	"sync"
)

// TaobaoWdkEquipmentConveyorSystemeventGetResult 结构体
type TaobaoWdkEquipmentConveyorSystemeventGetResult struct {
	// 返回的数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 返回值与返回的信息
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolTaobaoWdkEquipmentConveyorSystemeventGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorSystemeventGetResult)
	},
}

// GetTaobaoWdkEquipmentConveyorSystemeventGetResult() 从对象池中获取TaobaoWdkEquipmentConveyorSystemeventGetResult
func GetTaobaoWdkEquipmentConveyorSystemeventGetResult() *TaobaoWdkEquipmentConveyorSystemeventGetResult {
	return poolTaobaoWdkEquipmentConveyorSystemeventGetResult.Get().(*TaobaoWdkEquipmentConveyorSystemeventGetResult)
}

// ReleaseTaobaoWdkEquipmentConveyorSystemeventGetResult 释放TaobaoWdkEquipmentConveyorSystemeventGetResult
func ReleaseTaobaoWdkEquipmentConveyorSystemeventGetResult(v *TaobaoWdkEquipmentConveyorSystemeventGetResult) {
	v.Data = ""
	v.ResultCode = nil
	poolTaobaoWdkEquipmentConveyorSystemeventGetResult.Put(v)
}
