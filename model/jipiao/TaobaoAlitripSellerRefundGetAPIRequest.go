package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripsellerrefundgetAPIRequest 【机票代理商】退票申请单详情 API请求
// taobao.alitrip.seller.refund.get
//
// 查询退票申请单详情
type TaobaoalitripsellerrefundgetAPIRequest struct {
	model.Params
	// 申请单ID
	_applyId int64
}

// NewTaobaoalitripsellerrefundgetRequest 初始化TaobaoalitripsellerrefundgetAPIRequest对象
func NewTaobaoalitripsellerrefundgetRequest() *TaobaoalitripsellerrefundgetAPIRequest {
	return &TaobaoalitripsellerrefundgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripsellerrefundgetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripsellerrefundgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripsellerrefundgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单ID
func (r *TaobaoalitripsellerrefundgetAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoalitripsellerrefundgetAPIRequest) GetApplyId() int64 {
	return r._applyId
}
