package alihouse

import (
	"sync"
)

// QueryStorePunishDto 结构体
type QueryStorePunishDto struct {
	// 外部门店id列表
	OuterStoreIds []string `json:"outer_store_ids,omitempty" xml:"outer_store_ids>string,omitempty"`
}

var poolQueryStorePunishDto = sync.Pool{
	New: func() any {
		return new(QueryStorePunishDto)
	},
}

// GetQueryStorePunishDto() 从对象池中获取QueryStorePunishDto
func GetQueryStorePunishDto() *QueryStorePunishDto {
	return poolQueryStorePunishDto.Get().(*QueryStorePunishDto)
}

// ReleaseQueryStorePunishDto 释放QueryStorePunishDto
func ReleaseQueryStorePunishDto(v *QueryStorePunishDto) {
	v.OuterStoreIds = v.OuterStoreIds[:0]
	poolQueryStorePunishDto.Put(v)
}
