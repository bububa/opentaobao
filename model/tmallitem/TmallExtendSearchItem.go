package tmallitem

import (
	"sync"
)

// TmallExtendSearchItem 结构体
type TmallExtendSearchItem struct {
	// 邮费
	FastPostFee string `json:"fast_post_fee,omitempty" xml:"fast_post_fee,omitempty"`
	// 搜索宝贝的宝贝所在地
	Location string `json:"location,omitempty" xml:"location,omitempty"`
	// 搜索宝贝的卖家昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 搜索宝贝的图片url
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// 搜索宝贝的sku最低价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 搜索宝贝的sku最低价的折扣价
	PriceWithRate string `json:"price_with_rate,omitempty" xml:"price_with_rate,omitempty"`
	// 搜索宝贝的卖家所在地
	SellerLoc string `json:"seller_loc,omitempty" xml:"seller_loc,omitempty"`
	// 搜索宝贝的月数量
	Sold string `json:"sold,omitempty" xml:"sold,omitempty"`
	// 搜索宝贝的标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 搜索宝贝url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 是否货到付款
	IsCod int64 `json:"is_cod,omitempty" xml:"is_cod,omitempty"`
	// 搜索宝贝的数字id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 是否免邮
	Shipping int64 `json:"shipping,omitempty" xml:"shipping,omitempty"`
	// 搜索宝贝的spuid
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 搜索宝贝的卖家数字id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 是否折扣
	IsPromotion bool `json:"is_promotion,omitempty" xml:"is_promotion,omitempty"`
}

var poolTmallExtendSearchItem = sync.Pool{
	New: func() any {
		return new(TmallExtendSearchItem)
	},
}

// GetTmallExtendSearchItem() 从对象池中获取TmallExtendSearchItem
func GetTmallExtendSearchItem() *TmallExtendSearchItem {
	return poolTmallExtendSearchItem.Get().(*TmallExtendSearchItem)
}

// ReleaseTmallExtendSearchItem 释放TmallExtendSearchItem
func ReleaseTmallExtendSearchItem(v *TmallExtendSearchItem) {
	v.FastPostFee = ""
	v.Location = ""
	v.Nick = ""
	v.PicPath = ""
	v.Price = ""
	v.PriceWithRate = ""
	v.SellerLoc = ""
	v.Sold = ""
	v.Title = ""
	v.Url = ""
	v.IsCod = 0
	v.ItemId = 0
	v.Shipping = 0
	v.SpuId = 0
	v.UserId = 0
	v.IsPromotion = false
	poolTmallExtendSearchItem.Put(v)
}
