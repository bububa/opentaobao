package larkiot

import (
	"sync"
)

// ThirdGoodsListRsp 结构体
type ThirdGoodsListRsp struct {
	// 卖品列表
	GoodsList []ThirdGoodsRsp `json:"goods_list,omitempty" xml:"goods_list>third_goods_rsp,omitempty"`
	// 影院内码
	CinemaLinkId string `json:"cinema_link_id,omitempty" xml:"cinema_link_id,omitempty"`
	// 最大数量
	MaxBuyCount int64 `json:"max_buy_count,omitempty" xml:"max_buy_count,omitempty"`
}

var poolThirdGoodsListRsp = sync.Pool{
	New: func() any {
		return new(ThirdGoodsListRsp)
	},
}

// GetThirdGoodsListRsp() 从对象池中获取ThirdGoodsListRsp
func GetThirdGoodsListRsp() *ThirdGoodsListRsp {
	return poolThirdGoodsListRsp.Get().(*ThirdGoodsListRsp)
}

// ReleaseThirdGoodsListRsp 释放ThirdGoodsListRsp
func ReleaseThirdGoodsListRsp(v *ThirdGoodsListRsp) {
	v.GoodsList = v.GoodsList[:0]
	v.CinemaLinkId = ""
	v.MaxBuyCount = 0
	poolThirdGoodsListRsp.Put(v)
}
