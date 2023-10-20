package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkCouponGetAPIRequest 淘宝客-公用-阿里妈妈推广券详情查询 API请求
// taobao.tbk.coupon.get
//
// 传入商品ID+券ID(券ID已知情况下)，或者传入me参数，均可查询阿里妈妈推广券详细信息。
type TaobaoTbkCouponGetAPIRequest struct {
	model.Params
	// 带券ID与商品ID的加密串
	_me string
	// 商品ID
	_itemId string
	// 券ID
	_activityId string
}

// NewTaobaoTbkCouponGetRequest 初始化TaobaoTbkCouponGetAPIRequest对象
func NewTaobaoTbkCouponGetRequest() *TaobaoTbkCouponGetAPIRequest {
	return &TaobaoTbkCouponGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkCouponGetAPIRequest) Reset() {
	r._me = ""
	r._itemId = ""
	r._activityId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkCouponGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.coupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkCouponGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkCouponGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMe is Me Setter
// 带券ID与商品ID的加密串
func (r *TaobaoTbkCouponGetAPIRequest) SetMe(_me string) error {
	r._me = _me
	r.Set("me", _me)
	return nil
}

// GetMe Me Getter
func (r TaobaoTbkCouponGetAPIRequest) GetMe() string {
	return r._me
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoTbkCouponGetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoTbkCouponGetAPIRequest) GetItemId() string {
	return r._itemId
}

// SetActivityId is ActivityId Setter
// 券ID
func (r *TaobaoTbkCouponGetAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoTbkCouponGetAPIRequest) GetActivityId() string {
	return r._activityId
}

var poolTaobaoTbkCouponGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkCouponGetRequest()
	},
}

// GetTaobaoTbkCouponGetRequest 从 sync.Pool 获取 TaobaoTbkCouponGetAPIRequest
func GetTaobaoTbkCouponGetAPIRequest() *TaobaoTbkCouponGetAPIRequest {
	return poolTaobaoTbkCouponGetAPIRequest.Get().(*TaobaoTbkCouponGetAPIRequest)
}

// ReleaseTaobaoTbkCouponGetAPIRequest 将 TaobaoTbkCouponGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkCouponGetAPIRequest(v *TaobaoTbkCouponGetAPIRequest) {
	v.Reset()
	poolTaobaoTbkCouponGetAPIRequest.Put(v)
}
