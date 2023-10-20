package trade

import (
	"sync"
)

// IsvSimpleSubOrderModel 结构体
type IsvSimpleSubOrderModel struct {
	// 商品的类目(Key)，可不填写
	CargoKey string `json:"cargo_key,omitempty" xml:"cargo_key,omitempty"`
	// 1688的商品ID(offerId)
	OfferId string `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
	// 购买数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 1688的单品货号ID(skuId)，如果有的话，必须填写
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolIsvSimpleSubOrderModel = sync.Pool{
	New: func() any {
		return new(IsvSimpleSubOrderModel)
	},
}

// GetIsvSimpleSubOrderModel() 从对象池中获取IsvSimpleSubOrderModel
func GetIsvSimpleSubOrderModel() *IsvSimpleSubOrderModel {
	return poolIsvSimpleSubOrderModel.Get().(*IsvSimpleSubOrderModel)
}

// ReleaseIsvSimpleSubOrderModel 释放IsvSimpleSubOrderModel
func ReleaseIsvSimpleSubOrderModel(v *IsvSimpleSubOrderModel) {
	v.CargoKey = ""
	v.OfferId = ""
	v.Quantity = ""
	v.SkuId = ""
	poolIsvSimpleSubOrderModel.Put(v)
}
