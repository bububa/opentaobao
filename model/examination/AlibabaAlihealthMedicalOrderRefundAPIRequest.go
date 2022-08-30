package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalOrderRefundAPIRequest 退款接口 API请求
// alibaba.alihealth.medical.order.refund
//
// 退款接口
type AlibabaAlihealthMedicalOrderRefundAPIRequest struct {
	model.Params
	// 入参
	_orderRefundRequest *OrderRefundRequest
}

// NewAlibabaAlihealthMedicalOrderRefundRequest 初始化AlibabaAlihealthMedicalOrderRefundAPIRequest对象
func NewAlibabaAlihealthMedicalOrderRefundRequest() *AlibabaAlihealthMedicalOrderRefundAPIRequest {
	return &AlibabaAlihealthMedicalOrderRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalOrderRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalOrderRefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderRefundRequest is OrderRefundRequest Setter
// 入参
func (r *AlibabaAlihealthMedicalOrderRefundAPIRequest) SetOrderRefundRequest(_orderRefundRequest *OrderRefundRequest) error {
	r._orderRefundRequest = _orderRefundRequest
	r.Set("order_refund_request", _orderRefundRequest)
	return nil
}

// GetOrderRefundRequest OrderRefundRequest Getter
func (r AlibabaAlihealthMedicalOrderRefundAPIRequest) GetOrderRefundRequest() *OrderRefundRequest {
	return r._orderRefundRequest
}
