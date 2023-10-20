package wdk

import (
	"sync"
)

// TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult 结构体
type TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult struct {
	// 返回的数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 返回值与返回的信息
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolTaobaoWdkEquipmentConveyorHardwarestatuslogGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult)
	},
}

// GetTaobaoWdkEquipmentConveyorHardwarestatuslogGetResult() 从对象池中获取TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult
func GetTaobaoWdkEquipmentConveyorHardwarestatuslogGetResult() *TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult {
	return poolTaobaoWdkEquipmentConveyorHardwarestatuslogGetResult.Get().(*TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult)
}

// ReleaseTaobaoWdkEquipmentConveyorHardwarestatuslogGetResult 释放TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult
func ReleaseTaobaoWdkEquipmentConveyorHardwarestatuslogGetResult(v *TaobaoWdkEquipmentConveyorHardwarestatuslogGetResult) {
	v.Data = ""
	v.ResultCode = nil
	poolTaobaoWdkEquipmentConveyorHardwarestatuslogGetResult.Put(v)
}
