package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkskubestcouponAPIRequest sku维度最优优惠券信息 API请求
// taobao.tbk.sku.best.coupon
//
// 根据itemid和skuid查询最优优惠券信息
type TaobaotbkskubestcouponAPIRequest struct {
	model.Params
	// 商品对应的skuId
	_skuId int64
	// 商品ID
	_itemId int64
}

// NewTaobaotbkskubestcouponRequest 初始化TaobaotbkskubestcouponAPIRequest对象
func NewTaobaotbkskubestcouponRequest() *TaobaotbkskubestcouponAPIRequest {
	return &TaobaotbkskubestcouponAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkskubestcouponAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sku.best.coupon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkskubestcouponAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkskubestcouponAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuId is SkuId Setter
// 商品对应的skuId
func (r *TaobaotbkskubestcouponAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaotbkskubestcouponAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaotbkskubestcouponAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaotbkskubestcouponAPIRequest) GetItemId() int64 {
	return r._itemId
}
