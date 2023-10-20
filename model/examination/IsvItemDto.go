package examination

import (
	"sync"
)

// IsvItemDto 结构体
type IsvItemDto struct {
	// isv的单项id
	IsvItemId string `json:"isv_item_id,omitempty" xml:"isv_item_id,omitempty"`
	// 单项描述
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 售价，单位分
	SoldPrice string `json:"sold_price,omitempty" xml:"sold_price,omitempty"`
	// 版本号，防止isv更改未同步给健康，提供给isv做校验的
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolIsvItemDto = sync.Pool{
	New: func() any {
		return new(IsvItemDto)
	},
}

// GetIsvItemDto() 从对象池中获取IsvItemDto
func GetIsvItemDto() *IsvItemDto {
	return poolIsvItemDto.Get().(*IsvItemDto)
}

// ReleaseIsvItemDto 释放IsvItemDto
func ReleaseIsvItemDto(v *IsvItemDto) {
	v.IsvItemId = ""
	v.Detail = ""
	v.SubTitle = ""
	v.SoldPrice = ""
	v.Version = 0
	poolIsvItemDto.Put(v)
}
