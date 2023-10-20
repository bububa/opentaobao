package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripsellerrefundconfirmreturnAPIRequest 【机票代理商】确认退票 API请求
// taobao.alitrip.seller.refund.confirmreturn
//
// 确认退票
type TaobaoalitripsellerrefundconfirmreturnAPIRequest struct {
	model.Params
	// 退票申请单
	_applyId int64
}

// NewTaobaoalitripsellerrefundconfirmreturnRequest 初始化TaobaoalitripsellerrefundconfirmreturnAPIRequest对象
func NewTaobaoalitripsellerrefundconfirmreturnRequest() *TaobaoalitripsellerrefundconfirmreturnAPIRequest {
	return &TaobaoalitripsellerrefundconfirmreturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripsellerrefundconfirmreturnAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refund.confirmreturn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripsellerrefundconfirmreturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripsellerrefundconfirmreturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 退票申请单
func (r *TaobaoalitripsellerrefundconfirmreturnAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoalitripsellerrefundconfirmreturnAPIRequest) GetApplyId() int64 {
	return r._applyId
}
