package examination

import (
	"sync"
)

// IsvPackRelationDto 结构体
type IsvPackRelationDto struct {
	// 加项包id
	IsvPackId string `json:"isv_pack_id,omitempty" xml:"isv_pack_id,omitempty"`
	// 关联加项包id
	RelIsvPackId string `json:"rel_isv_pack_id,omitempty" xml:"rel_isv_pack_id,omitempty"`
	// 关系：1、互斥.
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolIsvPackRelationDto = sync.Pool{
	New: func() any {
		return new(IsvPackRelationDto)
	},
}

// GetIsvPackRelationDto() 从对象池中获取IsvPackRelationDto
func GetIsvPackRelationDto() *IsvPackRelationDto {
	return poolIsvPackRelationDto.Get().(*IsvPackRelationDto)
}

// ReleaseIsvPackRelationDto 释放IsvPackRelationDto
func ReleaseIsvPackRelationDto(v *IsvPackRelationDto) {
	v.IsvPackId = ""
	v.RelIsvPackId = ""
	v.Type = 0
	poolIsvPackRelationDto.Put(v)
}
