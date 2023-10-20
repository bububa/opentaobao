package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoregionpricemanageAPIRequest 编辑区域价格 API请求
// taobao.region.price.manage
//
// 编辑区域价格
type TaobaoregionpricemanageAPIRequest struct {
	model.Params
	// 列表
	_regionalPriceDtos []RegionalPriceDto
	// 商品id
	_itemId int64
	// 无sku传0
	_skuId int64
	// true:全量, false:增量
	_isFull bool
}

// NewTaobaoregionpricemanageRequest 初始化TaobaoregionpricemanageAPIRequest对象
func NewTaobaoregionpricemanageRequest() *TaobaoregionpricemanageAPIRequest {
	return &TaobaoregionpricemanageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoregionpricemanageAPIRequest) GetApiMethodName() string {
	return "taobao.region.price.manage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoregionpricemanageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoregionpricemanageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRegionalPriceDtos is RegionalPriceDtos Setter
// 列表
func (r *TaobaoregionpricemanageAPIRequest) SetRegionalPriceDtos(_regionalPriceDtos []RegionalPriceDto) error {
	r._regionalPriceDtos = _regionalPriceDtos
	r.Set("regional_price_dtos", _regionalPriceDtos)
	return nil
}

// GetRegionalPriceDtos RegionalPriceDtos Getter
func (r TaobaoregionpricemanageAPIRequest) GetRegionalPriceDtos() []RegionalPriceDto {
	return r._regionalPriceDtos
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoregionpricemanageAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoregionpricemanageAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// 无sku传0
func (r *TaobaoregionpricemanageAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoregionpricemanageAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetIsFull is IsFull Setter
// true:全量, false:增量
func (r *TaobaoregionpricemanageAPIRequest) SetIsFull(_isFull bool) error {
	r._isFull = _isFull
	r.Set("is_full", _isFull)
	return nil
}

// GetIsFull IsFull Getter
func (r TaobaoregionpricemanageAPIRequest) GetIsFull() bool {
	return r._isFull
}
