package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdoctorleshuiauditresultAPIRequest 乐税审核结果通知 API请求
// alibaba.alihealth.doctor.leshui.audit.result
//
// 乐税审核结果通知
type AlibabaalihealthdoctorleshuiauditresultAPIRequest struct {
	model.Params
	// 入参
	_param *PayTaxNoticeRequest
}

// NewAlibabaalihealthdoctorleshuiauditresultRequest 初始化AlibabaalihealthdoctorleshuiauditresultAPIRequest对象
func NewAlibabaalihealthdoctorleshuiauditresultRequest() *AlibabaalihealthdoctorleshuiauditresultAPIRequest {
	return &AlibabaalihealthdoctorleshuiauditresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdoctorleshuiauditresultAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.doctor.leshui.audit.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdoctorleshuiauditresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdoctorleshuiauditresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaalihealthdoctorleshuiauditresultAPIRequest) SetParam(_param *PayTaxNoticeRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaalihealthdoctorleshuiauditresultAPIRequest) GetParam() *PayTaxNoticeRequest {
	return r._param
}
