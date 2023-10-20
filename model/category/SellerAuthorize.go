package category

import (
	"sync"
)

// SellerAuthorize 结构体
type SellerAuthorize struct {
	// 商品类目结构
	ItemCats []ItemCat `json:"item_cats,omitempty" xml:"item_cats>item_cat,omitempty"`
	// 商品类目结构
	XinpinItemCats []ItemCat `json:"xinpin_item_cats,omitempty" xml:"xinpin_item_cats>item_cat,omitempty"`
	// 品牌
	Brands []Brand `json:"brands,omitempty" xml:"brands>brand,omitempty"`
}

var poolSellerAuthorize = sync.Pool{
	New: func() any {
		return new(SellerAuthorize)
	},
}

// GetSellerAuthorize() 从对象池中获取SellerAuthorize
func GetSellerAuthorize() *SellerAuthorize {
	return poolSellerAuthorize.Get().(*SellerAuthorize)
}

// ReleaseSellerAuthorize 释放SellerAuthorize
func ReleaseSellerAuthorize(v *SellerAuthorize) {
	v.ItemCats = v.ItemCats[:0]
	v.XinpinItemCats = v.XinpinItemCats[:0]
	v.Brands = v.Brands[:0]
	poolSellerAuthorize.Put(v)
}
