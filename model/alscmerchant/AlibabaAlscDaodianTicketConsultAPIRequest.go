package alscmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscDaodianTicketConsultAPIRequest 券码预览 API请求
// alibaba.alsc.daodian.ticket.consult
//
// 券码预览
type AlibabaAlscDaodianTicketConsultAPIRequest struct {
	model.Params
	// 券码咨询请求
	_consultRequest *TicketConsultTopRequest
}

// NewAlibabaAlscDaodianTicketConsultRequest 初始化AlibabaAlscDaodianTicketConsultAPIRequest对象
func NewAlibabaAlscDaodianTicketConsultRequest() *AlibabaAlscDaodianTicketConsultAPIRequest {
	return &AlibabaAlscDaodianTicketConsultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscDaodianTicketConsultAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.daodian.ticket.consult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscDaodianTicketConsultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscDaodianTicketConsultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsultRequest is ConsultRequest Setter
// 券码咨询请求
func (r *AlibabaAlscDaodianTicketConsultAPIRequest) SetConsultRequest(_consultRequest *TicketConsultTopRequest) error {
	r._consultRequest = _consultRequest
	r.Set("consult_request", _consultRequest)
	return nil
}

// GetConsultRequest ConsultRequest Getter
func (r AlibabaAlscDaodianTicketConsultAPIRequest) GetConsultRequest() *TicketConsultTopRequest {
	return r._consultRequest
}
