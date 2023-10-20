package product

import (
	"sync"
)

// ItemSearchResult 结构体
type ItemSearchResult struct {
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 商品链接
	ItemUrl string `json:"item_url,omitempty" xml:"item_url,omitempty"`
	// 商品主图
	ItemMainPic string `json:"item_main_pic,omitempty" xml:"item_main_pic,omitempty"`
	// 商品原价最低值
	ItemOriginPriceMin string `json:"item_origin_price_min,omitempty" xml:"item_origin_price_min,omitempty"`
	// 商品原价最高值
	ItemOriginPriceMax string `json:"item_origin_price_max,omitempty" xml:"item_origin_price_max,omitempty"`
	// 商品折扣价最低值
	ItemPriceDiscountMin string `json:"item_price_discount_min,omitempty" xml:"item_price_discount_min,omitempty"`
	// 商品折扣价最高值
	ItemPriceDiscountMax string `json:"item_price_discount_max,omitempty" xml:"item_price_discount_max,omitempty"`
	// 商品折扣率
	ItemDiscountRate string `json:"item_discount_rate,omitempty" xml:"item_discount_rate,omitempty"`
	// 商品发布时间
	PubTime string `json:"pub_time,omitempty" xml:"pub_time,omitempty"`
	// 商品评价分数
	CommentScore string `json:"comment_score,omitempty" xml:"comment_score,omitempty"`
	// 店铺URL
	ShopUrl string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
	// 佣金比例
	CommissionRate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// 商品图片地址
	ItemPics string `json:"item_pics,omitempty" xml:"item_pics,omitempty"`
	// 商品视频地址
	ItemVideos string `json:"item_videos,omitempty" xml:"item_videos,omitempty"`
	// 卖家服务等级
	SellerLayer string `json:"seller_layer,omitempty" xml:"seller_layer,omitempty"`
	// 30天销量语义化信息
	Sales30DaySemantic string `json:"sales30_day_semantic,omitempty" xml:"sales30_day_semantic,omitempty"`
	// 30天评论数语义化信息
	Comment30DaySemantic string `json:"comment30_day_semantic,omitempty" xml:"comment30_day_semantic,omitempty"`
	// 收藏数语义化信息
	FavoriteCntSemantic string `json:"favorite_cnt_semantic,omitempty" xml:"favorite_cnt_semantic,omitempty"`
}

var poolItemSearchResult = sync.Pool{
	New: func() any {
		return new(ItemSearchResult)
	},
}

// GetItemSearchResult() 从对象池中获取ItemSearchResult
func GetItemSearchResult() *ItemSearchResult {
	return poolItemSearchResult.Get().(*ItemSearchResult)
}

// ReleaseItemSearchResult 释放ItemSearchResult
func ReleaseItemSearchResult(v *ItemSearchResult) {
	v.ItemId = ""
	v.ItemTitle = ""
	v.ItemUrl = ""
	v.ItemMainPic = ""
	v.ItemOriginPriceMin = ""
	v.ItemOriginPriceMax = ""
	v.ItemPriceDiscountMin = ""
	v.ItemPriceDiscountMax = ""
	v.ItemDiscountRate = ""
	v.PubTime = ""
	v.CommentScore = ""
	v.ShopUrl = ""
	v.CommissionRate = ""
	v.ItemPics = ""
	v.ItemVideos = ""
	v.SellerLayer = ""
	v.Sales30DaySemantic = ""
	v.Comment30DaySemantic = ""
	v.FavoriteCntSemantic = ""
	poolItemSearchResult.Put(v)
}
