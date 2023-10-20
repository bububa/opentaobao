package opentrade

import (
	"sync"
)

// McUserDto 结构体
type McUserDto struct {
	// 用户状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 排队活动ID
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 排队传入的扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 用户openId
	UserOpenId string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
	// 排队商品数量
	Quality int64 `json:"quality,omitempty" xml:"quality,omitempty"`
	// 排队商品SKU ID
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 排队商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolMcUserDto = sync.Pool{
	New: func() any {
		return new(McUserDto)
	},
}

// GetMcUserDto() 从对象池中获取McUserDto
func GetMcUserDto() *McUserDto {
	return poolMcUserDto.Get().(*McUserDto)
}

// ReleaseMcUserDto 释放McUserDto
func ReleaseMcUserDto(v *McUserDto) {
	v.Status = ""
	v.ActivityId = ""
	v.ExtInfo = ""
	v.UserOpenId = ""
	v.Quality = 0
	v.SkuId = 0
	v.ItemId = 0
	poolMcUserDto.Put(v)
}
