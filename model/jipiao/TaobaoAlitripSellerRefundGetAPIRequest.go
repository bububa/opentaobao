package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundGetAPIRequest 【机票代理商】退票申请单详情 API请求
// taobao.alitrip.seller.refund.get
//
// 查询退票申请单详情
type TaobaoAlitripSellerRefundGetAPIRequest struct {
	model.Params
	// 申请单ID
	_applyId int64
}

// NewTaobaoAlitripSellerRefundGetRequest 初始化TaobaoAlitripSellerRefundGetAPIRequest对象
func NewTaobaoAlitripSellerRefundGetRequest() *TaobaoAlitripSellerRefundGetAPIRequest {
	return &TaobaoAlitripSellerRefundGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripSellerRefundGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单ID
func (r *TaobaoAlitripSellerRefundGetAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripSellerRefundGetAPIRequest) GetApplyId() int64 {
	return r._applyId
}
