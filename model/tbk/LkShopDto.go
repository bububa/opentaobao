package tbk

import (
	"sync"
)

// LkShopDto 结构体
type LkShopDto struct {
	// 店铺ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolLkShopDto = sync.Pool{
	New: func() any {
		return new(LkShopDto)
	},
}

// GetLkShopDto() 从对象池中获取LkShopDto
func GetLkShopDto() *LkShopDto {
	return poolLkShopDto.Get().(*LkShopDto)
}

// ReleaseLkShopDto 释放LkShopDto
func ReleaseLkShopDto(v *LkShopDto) {
	v.ShopId = ""
	poolLkShopDto.Put(v)
}
