package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionPriceCancleAPIRequest 取消区域价格 API请求
// taobao.region.price.cancle
//
// 取消区域价格
type TaobaoRegionPriceCancleAPIRequest struct {
	model.Params
	// 商品
	_itemId int64
	// 无sku传0
	_skuId int64
}

// NewTaobaoRegionPriceCancleRequest 初始化TaobaoRegionPriceCancleAPIRequest对象
func NewTaobaoRegionPriceCancleRequest() *TaobaoRegionPriceCancleAPIRequest {
	return &TaobaoRegionPriceCancleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRegionPriceCancleAPIRequest) GetApiMethodName() string {
	return "taobao.region.price.cancle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRegionPriceCancleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 商品
func (r *TaobaoRegionPriceCancleAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoRegionPriceCancleAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is SkuId Setter
// 无sku传0
func (r *TaobaoRegionPriceCancleAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r TaobaoRegionPriceCancleAPIRequest) GetSkuId() int64 {
	return r._skuId
}
