package examination

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalOrderRefundAPIRequest) Reset() {
	r._orderRefundRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalOrderRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalOrderRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalOrderRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthMedicalOrderRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalOrderRefundRequest()
	},
}

// GetAlibabaAlihealthMedicalOrderRefundRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalOrderRefundAPIRequest
func GetAlibabaAlihealthMedicalOrderRefundAPIRequest() *AlibabaAlihealthMedicalOrderRefundAPIRequest {
	return poolAlibabaAlihealthMedicalOrderRefundAPIRequest.Get().(*AlibabaAlihealthMedicalOrderRefundAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalOrderRefundAPIRequest 将 AlibabaAlihealthMedicalOrderRefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalOrderRefundAPIRequest(v *AlibabaAlihealthMedicalOrderRefundAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalOrderRefundAPIRequest.Put(v)
}
