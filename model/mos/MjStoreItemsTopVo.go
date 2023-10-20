package mos

import (
	"sync"
)

// MjStoreItemsTopVo 结构体
type MjStoreItemsTopVo struct {
	// 商品列表
	ItemList []MjItemTopVo `json:"item_list,omitempty" xml:"item_list>mj_item_top_vo,omitempty"`
	// 专柜名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// bucketId
	BucketId string `json:"bucket_id,omitempty" xml:"bucket_id,omitempty"`
	// 专柜id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolMjStoreItemsTopVo = sync.Pool{
	New: func() any {
		return new(MjStoreItemsTopVo)
	},
}

// GetMjStoreItemsTopVo() 从对象池中获取MjStoreItemsTopVo
func GetMjStoreItemsTopVo() *MjStoreItemsTopVo {
	return poolMjStoreItemsTopVo.Get().(*MjStoreItemsTopVo)
}

// ReleaseMjStoreItemsTopVo 释放MjStoreItemsTopVo
func ReleaseMjStoreItemsTopVo(v *MjStoreItemsTopVo) {
	v.ItemList = v.ItemList[:0]
	v.StoreName = ""
	v.Uuid = ""
	v.BucketId = ""
	v.StoreId = 0
	poolMjStoreItemsTopVo.Put(v)
}
