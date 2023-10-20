package wdk

import (
	"sync"
)

// SeasonVersionParam 结构体
type SeasonVersionParam struct {
	// 参与换挡的门店列表
	ShopIds []string `json:"shop_ids,omitempty" xml:"shop_ids>string,omitempty"`
	// 商家档期号
	OutSeasonId string `json:"out_season_id,omitempty" xml:"out_season_id,omitempty"`
}

var poolSeasonVersionParam = sync.Pool{
	New: func() any {
		return new(SeasonVersionParam)
	},
}

// GetSeasonVersionParam() 从对象池中获取SeasonVersionParam
func GetSeasonVersionParam() *SeasonVersionParam {
	return poolSeasonVersionParam.Get().(*SeasonVersionParam)
}

// ReleaseSeasonVersionParam 释放SeasonVersionParam
func ReleaseSeasonVersionParam(v *SeasonVersionParam) {
	v.ShopIds = v.ShopIds[:0]
	v.OutSeasonId = ""
	poolSeasonVersionParam.Put(v)
}
