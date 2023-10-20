package tbk

import (
	"sync"
)

// FavoritesInfo 结构体
type FavoritesInfo struct {
	// 选品库详细信息
	FavoritesList []FavoritesDetail `json:"favorites_list,omitempty" xml:"favorites_list>favorites_detail,omitempty"`
	// 选品库总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolFavoritesInfo = sync.Pool{
	New: func() any {
		return new(FavoritesInfo)
	},
}

// GetFavoritesInfo() 从对象池中获取FavoritesInfo
func GetFavoritesInfo() *FavoritesInfo {
	return poolFavoritesInfo.Get().(*FavoritesInfo)
}

// ReleaseFavoritesInfo 释放FavoritesInfo
func ReleaseFavoritesInfo(v *FavoritesInfo) {
	v.FavoritesList = v.FavoritesList[:0]
	v.TotalCount = 0
	poolFavoritesInfo.Put(v)
}
