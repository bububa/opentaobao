package elife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoElifeLifecardRefundAPIRequest 品牌惠卡券冲正退还 API请求
// taobao.elife.lifecard.refund
//
// 淘宝生活汇消费卡虚拟卡，线下冲正退货接口
type TaobaoElifeLifecardRefundAPIRequest struct {
	model.Params
	// 请求参数
	_refundRequest *RefundRequest
}

// NewTaobaoElifeLifecardRefundRequest 初始化TaobaoElifeLifecardRefundAPIRequest对象
func NewTaobaoElifeLifecardRefundRequest() *TaobaoElifeLifecardRefundAPIRequest {
	return &TaobaoElifeLifecardRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoElifeLifecardRefundAPIRequest) GetApiMethodName() string {
	return "taobao.elife.lifecard.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoElifeLifecardRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoElifeLifecardRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundRequest is RefundRequest Setter
// 请求参数
func (r *TaobaoElifeLifecardRefundAPIRequest) SetRefundRequest(_refundRequest *RefundRequest) error {
	r._refundRequest = _refundRequest
	r.Set("refund_request", _refundRequest)
	return nil
}

// GetRefundRequest RefundRequest Getter
func (r TaobaoElifeLifecardRefundAPIRequest) GetRefundRequest() *RefundRequest {
	return r._refundRequest
}
