package alihouse

import (
	"sync"
)

// ItemTagResultDto 结构体
type ItemTagResultDto struct {
	// 失败的外部ID列表
	FailedOuterIds []string `json:"failed_outer_ids,omitempty" xml:"failed_outer_ids>string,omitempty"`
	// 是否全部成功
	HasSuccess bool `json:"has_success,omitempty" xml:"has_success,omitempty"`
}

var poolItemTagResultDto = sync.Pool{
	New: func() any {
		return new(ItemTagResultDto)
	},
}

// GetItemTagResultDto() 从对象池中获取ItemTagResultDto
func GetItemTagResultDto() *ItemTagResultDto {
	return poolItemTagResultDto.Get().(*ItemTagResultDto)
}

// ReleaseItemTagResultDto 释放ItemTagResultDto
func ReleaseItemTagResultDto(v *ItemTagResultDto) {
	v.FailedOuterIds = v.FailedOuterIds[:0]
	v.HasSuccess = false
	poolItemTagResultDto.Put(v)
}
