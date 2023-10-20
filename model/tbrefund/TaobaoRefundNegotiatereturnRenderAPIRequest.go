package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundnegotiatereturnrenderAPIRequest 协商退货退款渲染 API请求
// taobao.refund.negotiatereturn.render
//
// 协商退货退款渲染
type TaobaorefundnegotiatereturnrenderAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
}

// NewTaobaorefundnegotiatereturnrenderRequest 初始化TaobaorefundnegotiatereturnrenderAPIRequest对象
func NewTaobaorefundnegotiatereturnrenderRequest() *TaobaorefundnegotiatereturnrenderAPIRequest {
	return &TaobaorefundnegotiatereturnrenderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorefundnegotiatereturnrenderAPIRequest) GetApiMethodName() string {
	return "taobao.refund.negotiatereturn.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorefundnegotiatereturnrenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorefundnegotiatereturnrenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaorefundnegotiatereturnrenderAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaorefundnegotiatereturnrenderAPIRequest) GetRefundId() int64 {
	return r._refundId
}
