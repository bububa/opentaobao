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
	_skuid int64
	// 商品ID
	_itemid int64
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

// SetSkuid is Skuid Setter
// 商品对应的skuId
func (r *TaobaotbkskubestcouponAPIRequest) SetSkuid(_skuid int64) error {
	r._skuid = _skuid
	r.Set("sku_id", _skuid)
	return nil
}

// GetSkuid Skuid Getter
func (r TaobaotbkskubestcouponAPIRequest) GetSkuid() int64 {
	return r._skuid
}

// SetItemid is Itemid Setter
// 商品ID
func (r *TaobaotbkskubestcouponAPIRequest) SetItemid(_itemid int64) error {
	r._itemid = _itemid
	r.Set("item_id", _itemid)
	return nil
}

// GetItemid Itemid Getter
func (r TaobaotbkskubestcouponAPIRequest) GetItemid() int64 {
	return r._itemid
}
