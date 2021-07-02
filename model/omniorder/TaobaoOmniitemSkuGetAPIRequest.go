package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemSkuGetAPIRequest 获取全渠道门店商品sku API请求
// taobao.omniitem.sku.get
//
// 通过skuId或者skuOutId查询全渠道门店商品sku信息
type TaobaoOmniitemSkuGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// skuId
	_skuId int64
	// sku商家编码
	_skuOuterId string
}

// NewTaobaoOmniitemSkuGetRequest 初始化TaobaoOmniitemSkuGetAPIRequest对象
func NewTaobaoOmniitemSkuGetRequest() *TaobaoOmniitemSkuGetAPIRequest {
	return &TaobaoOmniitemSkuGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemSkuGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemSkuGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 商品id
func (r *TaobaoOmniitemSkuGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoOmniitemSkuGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is SkuId Setter
// skuId
func (r *TaobaoOmniitemSkuGetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r TaobaoOmniitemSkuGetAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// Set is SkuOuterId Setter
// sku商家编码
func (r *TaobaoOmniitemSkuGetAPIRequest) SetSkuOuterId(_skuOuterId string) error {
	r._skuOuterId = _skuOuterId
	r.Set("sku_outer_id", _skuOuterId)
	return nil
}

// Get SkuOuterId Getter
func (r TaobaoOmniitemSkuGetAPIRequest) GetSkuOuterId() string {
	return r._skuOuterId
}
