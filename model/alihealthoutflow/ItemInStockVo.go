package alihealthoutflow

import (
	"sync"
)

// ItemInStockVo 结构体
type ItemInStockVo struct {
	// tab名
	TabName string `json:"tab_name,omitempty" xml:"tab_name,omitempty"`
	// 供应商地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 客户端名
	ClientName string `json:"client_name,omitempty" xml:"client_name,omitempty"`
	// 原价
	OriginPrice string `json:"origin_price,omitempty" xml:"origin_price,omitempty"`
	// 数据来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 商品名
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 标签名
	Tags string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 埋点
	EntryId string `json:"entry_id,omitempty" xml:"entry_id,omitempty"`
	// 商品图
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 现价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolItemInStockVo = sync.Pool{
	New: func() any {
		return new(ItemInStockVo)
	},
}

// GetItemInStockVo() 从对象池中获取ItemInStockVo
func GetItemInStockVo() *ItemInStockVo {
	return poolItemInStockVo.Get().(*ItemInStockVo)
}

// ReleaseItemInStockVo 释放ItemInStockVo
func ReleaseItemInStockVo(v *ItemInStockVo) {
	v.TabName = ""
	v.Address = ""
	v.ClientName = ""
	v.OriginPrice = ""
	v.Source = ""
	v.Title = ""
	v.Tags = ""
	v.EntryId = ""
	v.ImgUrl = ""
	v.Price = ""
	v.ItemId = 0
	poolItemInStockVo.Put(v)
}
