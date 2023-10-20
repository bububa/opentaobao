package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoregionpricequeryAPIRequest 区域价格查询 API请求
// taobao.region.price.query
//
// 区域价格查询
type TaobaoregionpricequeryAPIRequest struct {
	model.Params
	// 不传则返回所有设置的区域价格
	_regionalPriceDtos []RegionalPriceDto
	// 商品id
	_itemId int64
	// 无sku可传0
	_skuId int64
}

// NewTaobaoregionpricequeryRequest 初始化TaobaoregionpricequeryAPIRequest对象
func NewTaobaoregionpricequeryRequest() *TaobaoregionpricequeryAPIRequest {
	return &TaobaoregionpricequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoregionpricequeryAPIRequest) GetApiMethodName() string {
	return "taobao.region.price.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoregionpricequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoregionpricequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRegionalPriceDtos is RegionalPriceDtos Setter
// 不传则返回所有设置的区域价格
func (r *TaobaoregionpricequeryAPIRequest) SetRegionalPriceDtos(_regionalPriceDtos []RegionalPriceDto) error {
	r._regionalPriceDtos = _regionalPriceDtos
	r.Set("regional_price_dtos", _regionalPriceDtos)
	return nil
}

// GetRegionalPriceDtos RegionalPriceDtos Getter
func (r TaobaoregionpricequeryAPIRequest) GetRegionalPriceDtos() []RegionalPriceDto {
	return r._regionalPriceDtos
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoregionpricequeryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoregionpricequeryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// 无sku可传0
func (r *TaobaoregionpricequeryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoregionpricequeryAPIRequest) GetSkuId() int64 {
	return r._skuId
}
