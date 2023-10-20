package tmallhk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallHkClearanceCertificationGetAPIRequest 获取订单清关材料实名信息 API请求
// tmall.hk.clearance.certification.get
//
// 获取订单清关材料实名信息
type TmallHkClearanceCertificationGetAPIRequest struct {
	model.Params
	// 参数
	_orderCertRequest *OrderCertRequest
}

// NewTmallHkClearanceCertificationGetRequest 初始化TmallHkClearanceCertificationGetAPIRequest对象
func NewTmallHkClearanceCertificationGetRequest() *TmallHkClearanceCertificationGetAPIRequest {
	return &TmallHkClearanceCertificationGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallHkClearanceCertificationGetAPIRequest) Reset() {
	r._orderCertRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallHkClearanceCertificationGetAPIRequest) GetApiMethodName() string {
	return "tmall.hk.clearance.certification.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallHkClearanceCertificationGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallHkClearanceCertificationGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCertRequest is OrderCertRequest Setter
// 参数
func (r *TmallHkClearanceCertificationGetAPIRequest) SetOrderCertRequest(_orderCertRequest *OrderCertRequest) error {
	r._orderCertRequest = _orderCertRequest
	r.Set("order_cert_request", _orderCertRequest)
	return nil
}

// GetOrderCertRequest OrderCertRequest Getter
func (r TmallHkClearanceCertificationGetAPIRequest) GetOrderCertRequest() *OrderCertRequest {
	return r._orderCertRequest
}

var poolTmallHkClearanceCertificationGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallHkClearanceCertificationGetRequest()
	},
}

// GetTmallHkClearanceCertificationGetRequest 从 sync.Pool 获取 TmallHkClearanceCertificationGetAPIRequest
func GetTmallHkClearanceCertificationGetAPIRequest() *TmallHkClearanceCertificationGetAPIRequest {
	return poolTmallHkClearanceCertificationGetAPIRequest.Get().(*TmallHkClearanceCertificationGetAPIRequest)
}

// ReleaseTmallHkClearanceCertificationGetAPIRequest 将 TmallHkClearanceCertificationGetAPIRequest 放入 sync.Pool
func ReleaseTmallHkClearanceCertificationGetAPIRequest(v *TmallHkClearanceCertificationGetAPIRequest) {
	v.Reset()
	poolTmallHkClearanceCertificationGetAPIRequest.Put(v)
}
