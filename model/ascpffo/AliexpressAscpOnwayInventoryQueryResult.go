package ascpffo

import (
	"sync"
)

// AliexpressAscpOnwayInventoryQueryResult 结构体
type AliexpressAscpOnwayInventoryQueryResult struct {
	// 出参列表
	DataList []ErpOnWayInventoryDto `json:"data_list,omitempty" xml:"data_list>erp_on_way_inventory_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAliexpressAscpOnwayInventoryQueryResult = sync.Pool{
	New: func() any {
		return new(AliexpressAscpOnwayInventoryQueryResult)
	},
}

// GetAliexpressAscpOnwayInventoryQueryResult() 从对象池中获取AliexpressAscpOnwayInventoryQueryResult
func GetAliexpressAscpOnwayInventoryQueryResult() *AliexpressAscpOnwayInventoryQueryResult {
	return poolAliexpressAscpOnwayInventoryQueryResult.Get().(*AliexpressAscpOnwayInventoryQueryResult)
}

// ReleaseAliexpressAscpOnwayInventoryQueryResult 释放AliexpressAscpOnwayInventoryQueryResult
func ReleaseAliexpressAscpOnwayInventoryQueryResult(v *AliexpressAscpOnwayInventoryQueryResult) {
	v.DataList = v.DataList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAliexpressAscpOnwayInventoryQueryResult.Put(v)
}
