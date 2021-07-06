package tbk

import (
	"net/url"

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
	// 券ID
	_activityId string
	// 商品ID
	_itemId int64
}

// NewTaobaoTbkCouponGetRequest 初始化TaobaoTbkCouponGetAPIRequest对象
func NewTaobaoTbkCouponGetRequest() *TaobaoTbkCouponGetAPIRequest {
	return &TaobaoTbkCouponGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkCouponGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.coupon.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkCouponGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoTbkCouponGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoTbkCouponGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
