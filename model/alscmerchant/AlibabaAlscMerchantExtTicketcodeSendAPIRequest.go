package alscmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscmerchantextticketcodesendAPIRequest 异步发码 API请求
// alibaba.alsc.merchant.ext.ticketcode.send
//
// 外部券异步发码
type AlibabaalscmerchantextticketcodesendAPIRequest struct {
	model.Params
	// 外部券异步发码
	_sendRequest *ExternalTicketSendRequest
}

// NewAlibabaalscmerchantextticketcodesendRequest 初始化AlibabaalscmerchantextticketcodesendAPIRequest对象
func NewAlibabaalscmerchantextticketcodesendRequest() *AlibabaalscmerchantextticketcodesendAPIRequest {
	return &AlibabaalscmerchantextticketcodesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscmerchantextticketcodesendAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.merchant.ext.ticketcode.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscmerchantextticketcodesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscmerchantextticketcodesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSendRequest is SendRequest Setter
// 外部券异步发码
func (r *AlibabaalscmerchantextticketcodesendAPIRequest) SetSendRequest(_sendRequest *ExternalTicketSendRequest) error {
	r._sendRequest = _sendRequest
	r.Set("send_request", _sendRequest)
	return nil
}

// GetSendRequest SendRequest Getter
func (r AlibabaalscmerchantextticketcodesendAPIRequest) GetSendRequest() *ExternalTicketSendRequest {
	return r._sendRequest
}
