package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayRefundAPIRequest 退款接口 API请求
// aliyun.alios.pay.refund
//
// 商户用来发起退款的接口
type AliyunAliosPayRefundAPIRequest struct {
	model.Params
	// 请求参数
	_refundRequest *RefundRequest
}

// NewAliyunAliosPayRefundRequest 初始化AliyunAliosPayRefundAPIRequest对象
func NewAliyunAliosPayRefundRequest() *AliyunAliosPayRefundAPIRequest {
	return &AliyunAliosPayRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAliosPayRefundAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAliosPayRefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefundRequest is RefundRequest Setter
// 请求参数
func (r *AliyunAliosPayRefundAPIRequest) SetRefundRequest(_refundRequest *RefundRequest) error {
	r._refundRequest = _refundRequest
	r.Set("refund_request", _refundRequest)
	return nil
}

// GetRefundRequest RefundRequest Getter
func (r AliyunAliosPayRefundAPIRequest) GetRefundRequest() *RefundRequest {
	return r._refundRequest
}
