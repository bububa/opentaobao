package wdk

import (
	"sync"
)

// ItemBuyGiftSku 结构体
type ItemBuyGiftSku struct {
	// 商品的skuCode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 赠品的skuCode
	GiftSkuCode string `json:"gift_sku_code,omitempty" xml:"gift_sku_code,omitempty"`
	// 淘宝item和shop的对应关系， k-itemId, v-shopId
	ItemShopRelation string `json:"item_shop_relation,omitempty" xml:"item_shop_relation,omitempty"`
	// 主商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 赠品名称
	GiftSkuName string `json:"gift_sku_name,omitempty" xml:"gift_sku_name,omitempty"`
	// 限购信息
	LimitInfo *LimitInfo `json:"limit_info,omitempty" xml:"limit_info,omitempty"`
	// 买赠门槛数量
	BuyNum int64 `json:"buy_num,omitempty" xml:"buy_num,omitempty"`
}

var poolItemBuyGiftSku = sync.Pool{
	New: func() any {
		return new(ItemBuyGiftSku)
	},
}

// GetItemBuyGiftSku() 从对象池中获取ItemBuyGiftSku
func GetItemBuyGiftSku() *ItemBuyGiftSku {
	return poolItemBuyGiftSku.Get().(*ItemBuyGiftSku)
}

// ReleaseItemBuyGiftSku 释放ItemBuyGiftSku
func ReleaseItemBuyGiftSku(v *ItemBuyGiftSku) {
	v.SkuCode = ""
	v.GiftSkuCode = ""
	v.ItemShopRelation = ""
	v.SkuName = ""
	v.GiftSkuName = ""
	v.LimitInfo = nil
	v.BuyNum = 0
	poolItemBuyGiftSku.Put(v)
}
