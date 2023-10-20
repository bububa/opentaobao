package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniitemskugetAPIRequest 获取全渠道门店商品sku API请求
// taobao.omniitem.sku.get
//
// 通过skuId或者skuOutId查询全渠道门店商品sku信息
type TaobaoomniitemskugetAPIRequest struct {
	model.Params
	// sku商家编码
	_skuOuterId string
	// 商品id
	_itemId int64
	// skuId
	_skuId int64
}

// NewTaobaoomniitemskugetRequest 初始化TaobaoomniitemskugetAPIRequest对象
func NewTaobaoomniitemskugetRequest() *TaobaoomniitemskugetAPIRequest {
	return &TaobaoomniitemskugetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniitemskugetAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniitemskugetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniitemskugetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuOuterId is SkuOuterId Setter
// sku商家编码
func (r *TaobaoomniitemskugetAPIRequest) SetSkuOuterId(_skuOuterId string) error {
	r._skuOuterId = _skuOuterId
	r.Set("sku_outer_id", _skuOuterId)
	return nil
}

// GetSkuOuterId SkuOuterId Getter
func (r TaobaoomniitemskugetAPIRequest) GetSkuOuterId() string {
	return r._skuOuterId
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoomniitemskugetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoomniitemskugetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// skuId
func (r *TaobaoomniitemskugetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoomniitemskugetAPIRequest) GetSkuId() int64 {
	return r._skuId
}
