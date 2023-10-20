package mos

import (
	"sync"
)

// MjStoresTopVo 结构体
type MjStoresTopVo struct {
	// storeInfoList
	StoreInfoList []StoreInfo `json:"store_info_list,omitempty" xml:"store_info_list>store_info,omitempty"`
	// 外部门店号
	OutMallId string `json:"out_mall_id,omitempty" xml:"out_mall_id,omitempty"`
	// 默认选中的专柜，0为全部品牌
	StoreDefault int64 `json:"store_default,omitempty" xml:"store_default,omitempty"`
	// 版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 1：普通 2：刷脸
	ScreenType int64 `json:"screen_type,omitempty" xml:"screen_type,omitempty"`
	// 商场id
	MallId int64 `json:"mall_id,omitempty" xml:"mall_id,omitempty"`
}

var poolMjStoresTopVo = sync.Pool{
	New: func() any {
		return new(MjStoresTopVo)
	},
}

// GetMjStoresTopVo() 从对象池中获取MjStoresTopVo
func GetMjStoresTopVo() *MjStoresTopVo {
	return poolMjStoresTopVo.Get().(*MjStoresTopVo)
}

// ReleaseMjStoresTopVo 释放MjStoresTopVo
func ReleaseMjStoresTopVo(v *MjStoresTopVo) {
	v.StoreInfoList = v.StoreInfoList[:0]
	v.OutMallId = ""
	v.StoreDefault = 0
	v.Version = 0
	v.ScreenType = 0
	v.MallId = 0
	poolMjStoresTopVo.Put(v)
}
