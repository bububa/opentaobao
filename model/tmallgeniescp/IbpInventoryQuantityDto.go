package tmallgeniescp

import (
	"sync"
)

// IbpInventoryQuantityDto 结构体
type IbpInventoryQuantityDto struct {
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// 地点code
	LocationCode string `json:"location_code,omitempty" xml:"location_code,omitempty"`
	// 物料code
	MaterielCode string `json:"materiel_code,omitempty" xml:"materiel_code,omitempty"`
	// 库存量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolIbpInventoryQuantityDto = sync.Pool{
	New: func() any {
		return new(IbpInventoryQuantityDto)
	},
}

// GetIbpInventoryQuantityDto() 从对象池中获取IbpInventoryQuantityDto
func GetIbpInventoryQuantityDto() *IbpInventoryQuantityDto {
	return poolIbpInventoryQuantityDto.Get().(*IbpInventoryQuantityDto)
}

// ReleaseIbpInventoryQuantityDto 释放IbpInventoryQuantityDto
func ReleaseIbpInventoryQuantityDto(v *IbpInventoryQuantityDto) {
	v.ExtendJson = ""
	v.Tenant = ""
	v.LocationCode = ""
	v.MaterielCode = ""
	v.Quantity = 0
	poolIbpInventoryQuantityDto.Put(v)
}
