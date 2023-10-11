package tmallhk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
