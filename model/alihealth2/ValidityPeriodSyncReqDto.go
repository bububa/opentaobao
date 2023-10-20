package alihealth2

import (
	"sync"
)

// ValidityPeriodSyncReqDto 结构体
type ValidityPeriodSyncReqDto struct {
	// 请求明细
	Items []ValidityPeriodItem `json:"items,omitempty" xml:"items>validity_period_item,omitempty"`
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 供应商ID
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

var poolValidityPeriodSyncReqDto = sync.Pool{
	New: func() any {
		return new(ValidityPeriodSyncReqDto)
	},
}

// GetValidityPeriodSyncReqDto() 从对象池中获取ValidityPeriodSyncReqDto
func GetValidityPeriodSyncReqDto() *ValidityPeriodSyncReqDto {
	return poolValidityPeriodSyncReqDto.Get().(*ValidityPeriodSyncReqDto)
}

// ReleaseValidityPeriodSyncReqDto 释放ValidityPeriodSyncReqDto
func ReleaseValidityPeriodSyncReqDto(v *ValidityPeriodSyncReqDto) {
	v.Items = v.Items[:0]
	v.StoreCode = ""
	v.SupplierId = 0
	poolValidityPeriodSyncReqDto.Put(v)
}
