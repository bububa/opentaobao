package opentrade

import (
	"sync"
)

// ItemResultDto 结构体
type ItemResultDto struct {
	// 失败原因
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 失败商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolItemResultDto = sync.Pool{
	New: func() any {
		return new(ItemResultDto)
	},
}

// GetItemResultDto() 从对象池中获取ItemResultDto
func GetItemResultDto() *ItemResultDto {
	return poolItemResultDto.Get().(*ItemResultDto)
}

// ReleaseItemResultDto 释放ItemResultDto
func ReleaseItemResultDto(v *ItemResultDto) {
	v.ErrMsg = ""
	v.ItemId = 0
	poolItemResultDto.Put(v)
}
