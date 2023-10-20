package tbk

import (
	"sync"
)

// FavoritesDetail 结构体
type FavoritesDetail struct {
	// 选品库标题
	FavoritesTitle string `json:"favorites_title,omitempty" xml:"favorites_title,omitempty"`
	// 选品库id
	FavoritesId int64 `json:"favorites_id,omitempty" xml:"favorites_id,omitempty"`
}

var poolFavoritesDetail = sync.Pool{
	New: func() any {
		return new(FavoritesDetail)
	},
}

// GetFavoritesDetail() 从对象池中获取FavoritesDetail
func GetFavoritesDetail() *FavoritesDetail {
	return poolFavoritesDetail.Get().(*FavoritesDetail)
}

// ReleaseFavoritesDetail 释放FavoritesDetail
func ReleaseFavoritesDetail(v *FavoritesDetail) {
	v.FavoritesTitle = ""
	v.FavoritesId = 0
	poolFavoritesDetail.Put(v)
}
