package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallhkclearancecertificationgetAPIRequest 获取订单清关材料实名信息 API请求
// tmall.hk.clearance.certification.get
//
// 获取订单清关材料实名信息
type TmallhkclearancecertificationgetAPIRequest struct {
	model.Params
	// 参数
	_orderCertRequest *OrderCertRequest
}

// NewTmallhkclearancecertificationgetRequest 初始化TmallhkclearancecertificationgetAPIRequest对象
func NewTmallhkclearancecertificationgetRequest() *TmallhkclearancecertificationgetAPIRequest {
	return &TmallhkclearancecertificationgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallhkclearancecertificationgetAPIRequest) GetApiMethodName() string {
	return "tmall.hk.clearance.certification.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallhkclearancecertificationgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallhkclearancecertificationgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCertRequest is OrderCertRequest Setter
// 参数
func (r *TmallhkclearancecertificationgetAPIRequest) SetOrderCertRequest(_orderCertRequest *OrderCertRequest) error {
	r._orderCertRequest = _orderCertRequest
	r.Set("order_cert_request", _orderCertRequest)
	return nil
}

// GetOrderCertRequest OrderCertRequest Getter
func (r TmallhkclearancecertificationgetAPIRequest) GetOrderCertRequest() *OrderCertRequest {
	return r._orderCertRequest
}
