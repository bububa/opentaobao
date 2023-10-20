package alitripmerchant

import (
	"sync"
)

// FavoriteStatusVo 结构体
type FavoriteStatusVo struct {
	// 是否收藏
	FavoriteStatus bool `json:"favorite_status,omitempty" xml:"favorite_status,omitempty"`
}

var poolFavoriteStatusVo = sync.Pool{
	New: func() any {
		return new(FavoriteStatusVo)
	},
}

// GetFavoriteStatusVo() 从对象池中获取FavoriteStatusVo
func GetFavoriteStatusVo() *FavoriteStatusVo {
	return poolFavoriteStatusVo.Get().(*FavoriteStatusVo)
}

// ReleaseFavoriteStatusVo 释放FavoriteStatusVo
func ReleaseFavoriteStatusVo(v *FavoriteStatusVo) {
	v.FavoriteStatus = false
	poolFavoriteStatusVo.Put(v)
}
