package tbk

import (
	"sync"
)

// TargetItemDto 结构体
type TargetItemDto struct {
	// 页面内定坑商品ID，用于素材-坑位还原
	ItemIdList []string `json:"item_id_list,omitempty" xml:"item_id_list>string,omitempty"`
}

var poolTargetItemDto = sync.Pool{
	New: func() any {
		return new(TargetItemDto)
	},
}

// GetTargetItemDto() 从对象池中获取TargetItemDto
func GetTargetItemDto() *TargetItemDto {
	return poolTargetItemDto.Get().(*TargetItemDto)
}

// ReleaseTargetItemDto 释放TargetItemDto
func ReleaseTargetItemDto(v *TargetItemDto) {
	v.ItemIdList = v.ItemIdList[:0]
	poolTargetItemDto.Put(v)
}
