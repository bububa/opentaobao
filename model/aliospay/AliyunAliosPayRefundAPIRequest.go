package aliospay

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunAliosPayRefundAPIRequest) Reset() {
	r._refundRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAliosPayRefundAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAliosPayRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunAliosPayRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAliyunAliosPayRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunAliosPayRefundRequest()
	},
}

// GetAliyunAliosPayRefundRequest 从 sync.Pool 获取 AliyunAliosPayRefundAPIRequest
func GetAliyunAliosPayRefundAPIRequest() *AliyunAliosPayRefundAPIRequest {
	return poolAliyunAliosPayRefundAPIRequest.Get().(*AliyunAliosPayRefundAPIRequest)
}

// ReleaseAliyunAliosPayRefundAPIRequest 将 AliyunAliosPayRefundAPIRequest 放入 sync.Pool
func ReleaseAliyunAliosPayRefundAPIRequest(v *AliyunAliosPayRefundAPIRequest) {
	v.Reset()
	poolAliyunAliosPayRefundAPIRequest.Put(v)
}
