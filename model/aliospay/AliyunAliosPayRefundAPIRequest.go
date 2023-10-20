package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunaliospayrefundAPIRequest 退款接口 API请求
// aliyun.alios.pay.refund
//
// 商户用来发起退款的接口
type AliyunaliospayrefundAPIRequest struct {
	model.Params
	// 请求参数
	_refundRequest *RefundRequest
}

// NewAliyunaliospayrefundRequest 初始化AliyunaliospayrefundAPIRequest对象
func NewAliyunaliospayrefundRequest() *AliyunaliospayrefundAPIRequest {
	return &AliyunaliospayrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunaliospayrefundAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunaliospayrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunaliospayrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundRequest is RefundRequest Setter
// 请求参数
func (r *AliyunaliospayrefundAPIRequest) SetRefundRequest(_refundRequest *RefundRequest) error {
	r._refundRequest = _refundRequest
	r.Set("refund_request", _refundRequest)
	return nil
}

// GetRefundRequest RefundRequest Getter
func (r AliyunaliospayrefundAPIRequest) GetRefundRequest() *RefundRequest {
	return r._refundRequest
}
