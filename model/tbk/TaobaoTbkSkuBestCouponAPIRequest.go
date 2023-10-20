package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkSkuBestCouponAPIRequest sku维度最优优惠券信息 API请求
// taobao.tbk.sku.best.coupon
//
// 根据itemid和skuid查询最优优惠券信息
type TaobaoTbkSkuBestCouponAPIRequest struct {
	model.Params
	// 商品对应的skuId
	_skuId int64
	// 商品ID
	_itemId int64
}

// NewTaobaoTbkSkuBestCouponRequest 初始化TaobaoTbkSkuBestCouponAPIRequest对象
func NewTaobaoTbkSkuBestCouponRequest() *TaobaoTbkSkuBestCouponAPIRequest {
	return &TaobaoTbkSkuBestCouponAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkSkuBestCouponAPIRequest) Reset() {
	r._skuId = 0
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkSkuBestCouponAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sku.best.coupon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkSkuBestCouponAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkSkuBestCouponAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuId is SkuId Setter
// 商品对应的skuId
func (r *TaobaoTbkSkuBestCouponAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoTbkSkuBestCouponAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoTbkSkuBestCouponAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoTbkSkuBestCouponAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoTbkSkuBestCouponAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkSkuBestCouponRequest()
	},
}

// GetTaobaoTbkSkuBestCouponRequest 从 sync.Pool 获取 TaobaoTbkSkuBestCouponAPIRequest
func GetTaobaoTbkSkuBestCouponAPIRequest() *TaobaoTbkSkuBestCouponAPIRequest {
	return poolTaobaoTbkSkuBestCouponAPIRequest.Get().(*TaobaoTbkSkuBestCouponAPIRequest)
}

// ReleaseTaobaoTbkSkuBestCouponAPIRequest 将 TaobaoTbkSkuBestCouponAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkSkuBestCouponAPIRequest(v *TaobaoTbkSkuBestCouponAPIRequest) {
	v.Reset()
	poolTaobaoTbkSkuBestCouponAPIRequest.Put(v)
}
