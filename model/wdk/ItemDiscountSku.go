package wdk

import (
	"sync"
)

// ItemDiscountSku 结构体
type ItemDiscountSku struct {
	// 商品的skuCode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 淘宝item和shop的对应关系， k-itemId, v-shopId
	ItemShopRelation string `json:"item_shop_relation,omitempty" xml:"item_shop_relation,omitempty"`
	// 限购信息
	LimitInfo *LimitInfo `json:"limit_info,omitempty" xml:"limit_info,omitempty"`
	// &lt;优惠明细,分为单位&gt;优惠类型为[减价],则代表直降金额[如700,表示商品直降7元];优惠券类型为[一口价],则代表一口价[如700,表示商品一口价为7元];优惠券类型为[打折],则代表折扣[如700,表示打7折]
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 门槛数量，金额值单位为分
	ConditionNum int64 `json:"condition_num,omitempty" xml:"condition_num,omitempty"`
	// 门槛类型，2：累计消费金额，3：累计购买次数
	ConditionType int64 `json:"condition_type,omitempty" xml:"condition_type,omitempty"`
}

var poolItemDiscountSku = sync.Pool{
	New: func() any {
		return new(ItemDiscountSku)
	},
}

// GetItemDiscountSku() 从对象池中获取ItemDiscountSku
func GetItemDiscountSku() *ItemDiscountSku {
	return poolItemDiscountSku.Get().(*ItemDiscountSku)
}

// ReleaseItemDiscountSku 释放ItemDiscountSku
func ReleaseItemDiscountSku(v *ItemDiscountSku) {
	v.SkuCode = ""
	v.SkuName = ""
	v.ItemShopRelation = ""
	v.LimitInfo = nil
	v.Value = 0
	v.ConditionNum = 0
	v.ConditionType = 0
	poolItemDiscountSku.Put(v)
}
