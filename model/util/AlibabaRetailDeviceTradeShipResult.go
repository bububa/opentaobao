package util

import (
	"sync"
)

// AlibabaRetailDeviceTradeShipResult 结构体
type AlibabaRetailDeviceTradeShipResult struct {
	// errorInfos
	ErrorInfos []ErrorInfo `json:"error_infos,omitempty" xml:"error_infos>error_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaRetailDeviceTradeShipResult = sync.Pool{
	New: func() any {
		return new(AlibabaRetailDeviceTradeShipResult)
	},
}

// GetAlibabaRetailDeviceTradeShipResult() 从对象池中获取AlibabaRetailDeviceTradeShipResult
func GetAlibabaRetailDeviceTradeShipResult() *AlibabaRetailDeviceTradeShipResult {
	return poolAlibabaRetailDeviceTradeShipResult.Get().(*AlibabaRetailDeviceTradeShipResult)
}

// ReleaseAlibabaRetailDeviceTradeShipResult 释放AlibabaRetailDeviceTradeShipResult
func ReleaseAlibabaRetailDeviceTradeShipResult(v *AlibabaRetailDeviceTradeShipResult) {
	v.ErrorInfos = v.ErrorInfos[:0]
	v.Success = false
	poolAlibabaRetailDeviceTradeShipResult.Put(v)
}
