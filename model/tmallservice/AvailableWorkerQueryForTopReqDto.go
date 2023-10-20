package tmallservice

import (
	"sync"
)

// AvailableWorkerQueryForTopReqDto 结构体
type AvailableWorkerQueryForTopReqDto struct {
	// 服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务skuCode
	ServiceSkuCode string `json:"service_sku_code,omitempty" xml:"service_sku_code,omitempty"`
	// 网点code
	ServiceStoreCode string `json:"service_store_code,omitempty" xml:"service_store_code,omitempty"`
	// 服务容量查询开始日期。使用场景预约开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 服务容量查询结束日期。使用场景预约结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 地址编码
	AreaCode int64 `json:"area_code,omitempty" xml:"area_code,omitempty"`
}

var poolAvailableWorkerQueryForTopReqDto = sync.Pool{
	New: func() any {
		return new(AvailableWorkerQueryForTopReqDto)
	},
}

// GetAvailableWorkerQueryForTopReqDto() 从对象池中获取AvailableWorkerQueryForTopReqDto
func GetAvailableWorkerQueryForTopReqDto() *AvailableWorkerQueryForTopReqDto {
	return poolAvailableWorkerQueryForTopReqDto.Get().(*AvailableWorkerQueryForTopReqDto)
}

// ReleaseAvailableWorkerQueryForTopReqDto 释放AvailableWorkerQueryForTopReqDto
func ReleaseAvailableWorkerQueryForTopReqDto(v *AvailableWorkerQueryForTopReqDto) {
	v.ServiceCode = ""
	v.ServiceSkuCode = ""
	v.ServiceStoreCode = ""
	v.StartDate = ""
	v.EndDate = ""
	v.AreaCode = 0
	poolAvailableWorkerQueryForTopReqDto.Put(v)
}
