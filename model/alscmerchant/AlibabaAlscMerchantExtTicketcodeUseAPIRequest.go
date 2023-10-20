package alscmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscmerchantextticketcodeuseAPIRequest 外部核销服务 API请求
// alibaba.alsc.merchant.ext.ticketcode.use
//
// 外部核销服务
type AlibabaalscmerchantextticketcodeuseAPIRequest struct {
	model.Params
	// 外部券使用请求
	_useRequest *ExternalTicketUseRequest
}

// NewAlibabaalscmerchantextticketcodeuseRequest 初始化AlibabaalscmerchantextticketcodeuseAPIRequest对象
func NewAlibabaalscmerchantextticketcodeuseRequest() *AlibabaalscmerchantextticketcodeuseAPIRequest {
	return &AlibabaalscmerchantextticketcodeuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscmerchantextticketcodeuseAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.merchant.ext.ticketcode.use"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscmerchantextticketcodeuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscmerchantextticketcodeuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUseRequest is UseRequest Setter
// 外部券使用请求
func (r *AlibabaalscmerchantextticketcodeuseAPIRequest) SetUseRequest(_useRequest *ExternalTicketUseRequest) error {
	r._useRequest = _useRequest
	r.Set("use_request", _useRequest)
	return nil
}

// GetUseRequest UseRequest Getter
func (r AlibabaalscmerchantextticketcodeuseAPIRequest) GetUseRequest() *ExternalTicketUseRequest {
	return r._useRequest
}
