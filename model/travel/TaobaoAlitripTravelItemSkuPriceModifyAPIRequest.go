package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemSkuPriceModifyAPIRequest 【API3.0】日期级别日历价格库存修改，增量维护 API请求
// taobao.alitrip.travel.item.sku.price.modify
//
// 【API3.0】日期级别日历价格库存增量维护
type TaobaoAlitripTravelItemSkuPriceModifyAPIRequest struct {
	model.Params
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品日历价格库存套餐
	_skus []PontusTravelItemSkuInfo
}

// NewTaobaoAlitripTravelItemSkuPriceModifyRequest 初始化TaobaoAlitripTravelItemSkuPriceModifyAPIRequest对象
func NewTaobaoAlitripTravelItemSkuPriceModifyRequest() *TaobaoAlitripTravelItemSkuPriceModifyAPIRequest {
	return &TaobaoAlitripTravelItemSkuPriceModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.sku.price.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// Get OutProductId Getter
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// Set is Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) SetSkus(_skus []PontusTravelItemSkuInfo) error {
	r._skus = _skus
	r.Set("skus", _skus)
	return nil
}

// Get Skus Getter
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetSkus() []PontusTravelItemSkuInfo {
	return r._skus
}
