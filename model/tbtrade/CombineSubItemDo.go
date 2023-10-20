package tbtrade

import (
	"sync"
)

// CombineSubItemDo 结构体
type CombineSubItemDo struct {
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商家外部编码(可与商家外部系统对接)。
	OuterIid string `json:"outer_iid,omitempty" xml:"outer_iid,omitempty"`
	// SKU标题
	SkuTitle string `json:"sku_title,omitempty" xml:"sku_title,omitempty"`
	// 套餐购成分品预计承诺时效(如果为时间格式为yyyy-MM-dd 则为绝对时间，为数字则为相对时间，比如7天内发货)
	EstconTime string `json:"estcon_time,omitempty" xml:"estcon_time,omitempty"`
	// outer_sku_id
	OuterSkuId string `json:"outer_sku_id,omitempty" xml:"outer_sku_id,omitempty"`
	// 商品数字编号
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品的最小库存单位Sku的id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 成分品原价
	OriginFee int64 `json:"origin_fee,omitempty" xml:"origin_fee,omitempty"`
	// 成分品套餐原价
	CombineSubItemFee int64 `json:"combine_sub_item_fee,omitempty" xml:"combine_sub_item_fee,omitempty"`
	// 教育优惠原价
	EduOriginalFee int64 `json:"edu_original_fee,omitempty" xml:"edu_original_fee,omitempty"`
	// 套餐购是否商品主子成分品��
	Ismain bool `json:"ismain,omitempty" xml:"ismain,omitempty"`
}

var poolCombineSubItemDo = sync.Pool{
	New: func() any {
		return new(CombineSubItemDo)
	},
}

// GetCombineSubItemDo() 从对象池中获取CombineSubItemDo
func GetCombineSubItemDo() *CombineSubItemDo {
	return poolCombineSubItemDo.Get().(*CombineSubItemDo)
}

// ReleaseCombineSubItemDo 释放CombineSubItemDo
func ReleaseCombineSubItemDo(v *CombineSubItemDo) {
	v.ItemName = ""
	v.OuterIid = ""
	v.SkuTitle = ""
	v.EstconTime = ""
	v.OuterSkuId = ""
	v.ItemId = 0
	v.SkuId = 0
	v.Quantity = 0
	v.OriginFee = 0
	v.CombineSubItemFee = 0
	v.EduOriginalFee = 0
	v.Ismain = false
	poolCombineSubItemDo.Put(v)
}
