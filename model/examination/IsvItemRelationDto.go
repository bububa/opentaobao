package examination

import (
	"sync"
)

// IsvItemRelationDto 结构体
type IsvItemRelationDto struct {
	// 单项id
	IsvItemId string `json:"isv_item_id,omitempty" xml:"isv_item_id,omitempty"`
	// 关联单项id
	RelIsvItemId string `json:"rel_isv_item_id,omitempty" xml:"rel_isv_item_id,omitempty"`
	// 关系：1、互斥 2、重复
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolIsvItemRelationDto = sync.Pool{
	New: func() any {
		return new(IsvItemRelationDto)
	},
}

// GetIsvItemRelationDto() 从对象池中获取IsvItemRelationDto
func GetIsvItemRelationDto() *IsvItemRelationDto {
	return poolIsvItemRelationDto.Get().(*IsvItemRelationDto)
}

// ReleaseIsvItemRelationDto 释放IsvItemRelationDto
func ReleaseIsvItemRelationDto(v *IsvItemRelationDto) {
	v.IsvItemId = ""
	v.RelIsvItemId = ""
	v.Type = 0
	poolIsvItemRelationDto.Put(v)
}
