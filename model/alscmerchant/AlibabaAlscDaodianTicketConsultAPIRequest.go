package alscmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscdaodianticketconsultAPIRequest 券码预览 API请求
// alibaba.alsc.daodian.ticket.consult
//
// 券码预览
type AlibabaalscdaodianticketconsultAPIRequest struct {
	model.Params
	// 券码咨询请求
	_consultRequest *TicketConsultTopRequest
}

// NewAlibabaalscdaodianticketconsultRequest 初始化AlibabaalscdaodianticketconsultAPIRequest对象
func NewAlibabaalscdaodianticketconsultRequest() *AlibabaalscdaodianticketconsultAPIRequest {
	return &AlibabaalscdaodianticketconsultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscdaodianticketconsultAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.daodian.ticket.consult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscdaodianticketconsultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscdaodianticketconsultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsultRequest is ConsultRequest Setter
// 券码咨询请求
func (r *AlibabaalscdaodianticketconsultAPIRequest) SetConsultRequest(_consultRequest *TicketConsultTopRequest) error {
	r._consultRequest = _consultRequest
	r.Set("consult_request", _consultRequest)
	return nil
}

// GetConsultRequest ConsultRequest Getter
func (r AlibabaalscdaodianticketconsultAPIRequest) GetConsultRequest() *TicketConsultTopRequest {
	return r._consultRequest
}
