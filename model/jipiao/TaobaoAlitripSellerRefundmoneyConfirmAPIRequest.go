package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripsellerrefundmoneyconfirmAPIRequest 【机票代理商订单】确认退款 API请求
// taobao.alitrip.seller.refundmoney.confirm
//
// 代理人确认退票申请单的退款
type TaobaoalitripsellerrefundmoneyconfirmAPIRequest struct {
	model.Params
	// 申请单id
	_applyId int64
}

// NewTaobaoalitripsellerrefundmoneyconfirmRequest 初始化TaobaoalitripsellerrefundmoneyconfirmAPIRequest对象
func NewTaobaoalitripsellerrefundmoneyconfirmRequest() *TaobaoalitripsellerrefundmoneyconfirmAPIRequest {
	return &TaobaoalitripsellerrefundmoneyconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripsellerrefundmoneyconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refundmoney.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripsellerrefundmoneyconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripsellerrefundmoneyconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单id
func (r *TaobaoalitripsellerrefundmoneyconfirmAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoalitripsellerrefundmoneyconfirmAPIRequest) GetApplyId() int64 {
	return r._applyId
}
