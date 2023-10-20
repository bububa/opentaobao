package nrt

import (
	"sync"
)

// TmallNrtStoreItemFromOnlineItemResultDto 结构体
type TmallNrtStoreItemFromOnlineItemResultDto struct {
	// 商品集合
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 接口执行错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallNrtStoreItemFromOnlineItemResultDto = sync.Pool{
	New: func() any {
		return new(TmallNrtStoreItemFromOnlineItemResultDto)
	},
}

// GetTmallNrtStoreItemFromOnlineItemResultDto() 从对象池中获取TmallNrtStoreItemFromOnlineItemResultDto
func GetTmallNrtStoreItemFromOnlineItemResultDto() *TmallNrtStoreItemFromOnlineItemResultDto {
	return poolTmallNrtStoreItemFromOnlineItemResultDto.Get().(*TmallNrtStoreItemFromOnlineItemResultDto)
}

// ReleaseTmallNrtStoreItemFromOnlineItemResultDto 释放TmallNrtStoreItemFromOnlineItemResultDto
func ReleaseTmallNrtStoreItemFromOnlineItemResultDto(v *TmallNrtStoreItemFromOnlineItemResultDto) {
	v.ItemIds = v.ItemIds[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolTmallNrtStoreItemFromOnlineItemResultDto.Put(v)
}
