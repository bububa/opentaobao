package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdoctorleshuiticketvalidAPIRequest 乐税token验证 API请求
// alibaba.alihealth.doctor.leshui.ticket.valid
//
// 乐税token验证
type AlibabaalihealthdoctorleshuiticketvalidAPIRequest struct {
	model.Params
	// 入参
	_param *PayTaxValidationRequest
}

// NewAlibabaalihealthdoctorleshuiticketvalidRequest 初始化AlibabaalihealthdoctorleshuiticketvalidAPIRequest对象
func NewAlibabaalihealthdoctorleshuiticketvalidRequest() *AlibabaalihealthdoctorleshuiticketvalidAPIRequest {
	return &AlibabaalihealthdoctorleshuiticketvalidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdoctorleshuiticketvalidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.doctor.leshui.ticket.valid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdoctorleshuiticketvalidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdoctorleshuiticketvalidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaalihealthdoctorleshuiticketvalidAPIRequest) SetParam(_param *PayTaxValidationRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaalihealthdoctorleshuiticketvalidAPIRequest) GetParam() *PayTaxValidationRequest {
	return r._param
}
