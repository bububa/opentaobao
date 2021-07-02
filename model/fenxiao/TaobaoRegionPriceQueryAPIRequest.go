package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionPriceQueryAPIRequest 区域价格查询 API请求
// taobao.region.price.query
//
// 区域价格查询
type TaobaoRegionPriceQueryAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 无sku可传0
	_skuId int64
	// 不传则返回所有设置的区域价格
	_regionalPriceDtos []RegionalPriceDto
}

// NewTaobaoRegionPriceQueryRequest 初始化TaobaoRegionPriceQueryAPIRequest对象
func NewTaobaoRegionPriceQueryRequest() *TaobaoRegionPriceQueryAPIRequest {
	return &TaobaoRegionPriceQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRegionPriceQueryAPIRequest) GetApiMethodName() string {
	return "taobao.region.price.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRegionPriceQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 商品id
func (r *TaobaoRegionPriceQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoRegionPriceQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is SkuId Setter
// 无sku可传0
func (r *TaobaoRegionPriceQueryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r TaobaoRegionPriceQueryAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// Set is RegionalPriceDtos Setter
// 不传则返回所有设置的区域价格
func (r *TaobaoRegionPriceQueryAPIRequest) SetRegionalPriceDtos(_regionalPriceDtos []RegionalPriceDto) error {
	r._regionalPriceDtos = _regionalPriceDtos
	r.Set("regional_price_dtos", _regionalPriceDtos)
	return nil
}

// Get RegionalPriceDtos Getter
func (r TaobaoRegionPriceQueryAPIRequest) GetRegionalPriceDtos() []RegionalPriceDto {
	return r._regionalPriceDtos
}
