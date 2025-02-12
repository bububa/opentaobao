package tbk

import (
	"sync"
)

// LkPageDto 结构体
type LkPageDto struct {
	// 页面内定坑商品ID，用于素材-坑位还原
	TargetItemList []string `json:"target_item_list,omitempty" xml:"target_item_list>string,omitempty"`
	// 会场ID
	PageId string `json:"page_id,omitempty" xml:"page_id,omitempty"`
}

var poolLkPageDto = sync.Pool{
	New: func() any {
		return new(LkPageDto)
	},
}

// GetLkPageDto() 从对象池中获取LkPageDto
func GetLkPageDto() *LkPageDto {
	return poolLkPageDto.Get().(*LkPageDto)
}

// ReleaseLkPageDto 释放LkPageDto
func ReleaseLkPageDto(v *LkPageDto) {
	v.TargetItemList = v.TargetItemList[:0]
	v.PageId = ""
	poolLkPageDto.Put(v)
}
