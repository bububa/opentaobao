package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoregionpricecancleAPIRequest 取消区域价格 API请求
// taobao.region.price.cancle
//
// 取消区域价格
type TaobaoregionpricecancleAPIRequest struct {
	model.Params
	// 商品
	_itemId int64
	// 无sku传0
	_skuId int64
}

// NewTaobaoregionpricecancleRequest 初始化TaobaoregionpricecancleAPIRequest对象
func NewTaobaoregionpricecancleRequest() *TaobaoregionpricecancleAPIRequest {
	return &TaobaoregionpricecancleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoregionpricecancleAPIRequest) GetApiMethodName() string {
	return "taobao.region.price.cancle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoregionpricecancleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoregionpricecancleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品
func (r *TaobaoregionpricecancleAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoregionpricecancleAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// 无sku传0
func (r *TaobaoregionpricecancleAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoregionpricecancleAPIRequest) GetSkuId() int64 {
	return r._skuId
}
