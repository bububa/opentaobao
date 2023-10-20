package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdoctorleshuiapplynotifyAPIRequest 申请单审核结果通知 API请求
// alibaba.alihealth.doctor.leshui.apply.notify
//
// 申请单审核结果通知
type AlibabaalihealthdoctorleshuiapplynotifyAPIRequest struct {
	model.Params
	// 入参
	_param *ApplyTaxNoticeRequest
}

// NewAlibabaalihealthdoctorleshuiapplynotifyRequest 初始化AlibabaalihealthdoctorleshuiapplynotifyAPIRequest对象
func NewAlibabaalihealthdoctorleshuiapplynotifyRequest() *AlibabaalihealthdoctorleshuiapplynotifyAPIRequest {
	return &AlibabaalihealthdoctorleshuiapplynotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdoctorleshuiapplynotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.doctor.leshui.apply.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdoctorleshuiapplynotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdoctorleshuiapplynotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaalihealthdoctorleshuiapplynotifyAPIRequest) SetParam(_param *ApplyTaxNoticeRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaalihealthdoctorleshuiapplynotifyAPIRequest) GetParam() *ApplyTaxNoticeRequest {
	return r._param
}
