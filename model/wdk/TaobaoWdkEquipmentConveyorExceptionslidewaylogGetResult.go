package wdk

import (
	"sync"
)

// TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult 结构体
type TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult struct {
	// 返回的数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 返回值与返回的信息
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolTaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult)
	},
}

// GetTaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult() 从对象池中获取TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult
func GetTaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult() *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult {
	return poolTaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult.Get().(*TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult)
}

// ReleaseTaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult 释放TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult
func ReleaseTaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult(v *TaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult) {
	v.Data = ""
	v.ResultCode = nil
	poolTaobaoWdkEquipmentConveyorExceptionslidewaylogGetResult.Put(v)
}
