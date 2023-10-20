package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripsellerrefundrefusereturnAPIRequest 【机票代理商】拒绝退票 API请求
// taobao.alitrip.seller.refund.refusereturn
//
// 拒绝退票
type TaobaoalitripsellerrefundrefusereturnAPIRequest struct {
	model.Params
	// 拒绝理由
	_reason string
	// 申请单ID
	_applyId int64
}

// NewTaobaoalitripsellerrefundrefusereturnRequest 初始化TaobaoalitripsellerrefundrefusereturnAPIRequest对象
func NewTaobaoalitripsellerrefundrefusereturnRequest() *TaobaoalitripsellerrefundrefusereturnAPIRequest {
	return &TaobaoalitripsellerrefundrefusereturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripsellerrefundrefusereturnAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refund.refusereturn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripsellerrefundrefusereturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripsellerrefundrefusereturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReason is Reason Setter
// 拒绝理由
func (r *TaobaoalitripsellerrefundrefusereturnAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoalitripsellerrefundrefusereturnAPIRequest) GetReason() string {
	return r._reason
}

// SetApplyId is ApplyId Setter
// 申请单ID
func (r *TaobaoalitripsellerrefundrefusereturnAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoalitripsellerrefundrefusereturnAPIRequest) GetApplyId() int64 {
	return r._applyId
}
