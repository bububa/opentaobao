package simba

import (
	"sync"
)

// ItemVo 结构体
type ItemVo struct {
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品图片
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 商品地址
	LinkUrl string `json:"link_url,omitempty" xml:"link_url,omitempty"`
	// 首次上架时间
	FirstStartsTime string `json:"first_starts_time,omitempty" xml:"first_starts_time,omitempty"`
	// 最近上架时间
	Starts string `json:"starts,omitempty" xml:"starts,omitempty"`
	// 主站类目路径，空格分隔多级
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 销量
	BidCount int64 `json:"bid_count,omitempty" xml:"bid_count,omitempty"`
	// 库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 是否定时上架,true:是,false:否
	Timing bool `json:"timing,omitempty" xml:"timing,omitempty"`
}

var poolItemVo = sync.Pool{
	New: func() any {
		return new(ItemVo)
	},
}

// GetItemVo() 从对象池中获取ItemVo
func GetItemVo() *ItemVo {
	return poolItemVo.Get().(*ItemVo)
}

// ReleaseItemVo 释放ItemVo
func ReleaseItemVo(v *ItemVo) {
	v.Title = ""
	v.ImgUrl = ""
	v.LinkUrl = ""
	v.FirstStartsTime = ""
	v.Starts = ""
	v.CategoryId = ""
	v.ItemId = 0
	v.Price = 0
	v.BidCount = 0
	v.Quantity = 0
	v.Timing = false
	poolItemVo.Put(v)
}
