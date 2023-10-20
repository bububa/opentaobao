package inventory

import (
	"sync"
)

// LocationRelationDto 结构体
type LocationRelationDto struct {
	// 实体code
	TargetStoreCode string `json:"target_store_code,omitempty" xml:"target_store_code,omitempty"`
	// 实体code
	SourceStoreCode string `json:"source_store_code,omitempty" xml:"source_store_code,omitempty"`
	// 状态  0 正常  -1 删除
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 实体类型 2：仓库 6：门店
	TargetInvStoreType int64 `json:"target_inv_store_type,omitempty" xml:"target_inv_store_type,omitempty"`
	// 实体类型 2：仓库 6：门店
	SourceInvStoreType int64 `json:"source_inv_store_type,omitempty" xml:"source_inv_store_type,omitempty"`
}

var poolLocationRelationDto = sync.Pool{
	New: func() any {
		return new(LocationRelationDto)
	},
}

// GetLocationRelationDto() 从对象池中获取LocationRelationDto
func GetLocationRelationDto() *LocationRelationDto {
	return poolLocationRelationDto.Get().(*LocationRelationDto)
}

// ReleaseLocationRelationDto 释放LocationRelationDto
func ReleaseLocationRelationDto(v *LocationRelationDto) {
	v.TargetStoreCode = ""
	v.SourceStoreCode = ""
	v.Status = 0
	v.TargetInvStoreType = 0
	v.SourceInvStoreType = 0
	poolLocationRelationDto.Put(v)
}
