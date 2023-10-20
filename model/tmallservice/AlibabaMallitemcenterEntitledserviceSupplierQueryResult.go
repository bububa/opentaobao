package tmallservice

import (
	"sync"
)

// AlibabaMallitemcenterEntitledserviceSupplierQueryResult 结构体
type AlibabaMallitemcenterEntitledserviceSupplierQueryResult struct {
	// 接口返回model
	ResultData *ResultData `json:"result_data,omitempty" xml:"result_data,omitempty"`
}

var poolAlibabaMallitemcenterEntitledserviceSupplierQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaMallitemcenterEntitledserviceSupplierQueryResult)
	},
}

// GetAlibabaMallitemcenterEntitledserviceSupplierQueryResult() 从对象池中获取AlibabaMallitemcenterEntitledserviceSupplierQueryResult
func GetAlibabaMallitemcenterEntitledserviceSupplierQueryResult() *AlibabaMallitemcenterEntitledserviceSupplierQueryResult {
	return poolAlibabaMallitemcenterEntitledserviceSupplierQueryResult.Get().(*AlibabaMallitemcenterEntitledserviceSupplierQueryResult)
}

// ReleaseAlibabaMallitemcenterEntitledserviceSupplierQueryResult 释放AlibabaMallitemcenterEntitledserviceSupplierQueryResult
func ReleaseAlibabaMallitemcenterEntitledserviceSupplierQueryResult(v *AlibabaMallitemcenterEntitledserviceSupplierQueryResult) {
	v.ResultData = nil
	poolAlibabaMallitemcenterEntitledserviceSupplierQueryResult.Put(v)
}
