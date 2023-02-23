package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest 乐税token验证 API请求
// alibaba.alihealth.doctor.leshui.ticket.valid
//
// 乐税token验证
type AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest struct {
	model.Params
	// 入参
	_param *PayTaxValidationRequest
}

// NewAlibabaAlihealthDoctorLeshuiTicketValidRequest 初始化AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest对象
func NewAlibabaAlihealthDoctorLeshuiTicketValidRequest() *AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest {
	return &AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.doctor.leshui.ticket.valid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest) SetParam(_param *PayTaxValidationRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest) GetParam() *PayTaxValidationRequest {
	return r._param
}
