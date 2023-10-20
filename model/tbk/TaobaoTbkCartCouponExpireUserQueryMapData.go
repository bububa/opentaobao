package tbk

import (
	"sync"
)

// TaobaoTbkCartCouponExpireUserQueryMapData 结构体
type TaobaoTbkCartCouponExpireUserQueryMapData struct {
	// 商品ID对应的sku集合
	SkuIdList []int64 `json:"sku_id_list,omitempty" xml:"sku_id_list>int64,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolTaobaoTbkCartCouponExpireUserQueryMapData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkCartCouponExpireUserQueryMapData)
	},
}

// GetTaobaoTbkCartCouponExpireUserQueryMapData() 从对象池中获取TaobaoTbkCartCouponExpireUserQueryMapData
func GetTaobaoTbkCartCouponExpireUserQueryMapData() *TaobaoTbkCartCouponExpireUserQueryMapData {
	return poolTaobaoTbkCartCouponExpireUserQueryMapData.Get().(*TaobaoTbkCartCouponExpireUserQueryMapData)
}

// ReleaseTaobaoTbkCartCouponExpireUserQueryMapData 释放TaobaoTbkCartCouponExpireUserQueryMapData
func ReleaseTaobaoTbkCartCouponExpireUserQueryMapData(v *TaobaoTbkCartCouponExpireUserQueryMapData) {
	v.SkuIdList = v.SkuIdList[:0]
	v.ItemId = 0
	poolTaobaoTbkCartCouponExpireUserQueryMapData.Put(v)
}
