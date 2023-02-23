package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest 申请单审核结果通知 API请求
// alibaba.alihealth.doctor.leshui.apply.notify
//
// 申请单审核结果通知
type AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest struct {
	model.Params
	// 入参
	_param *ApplyTaxNoticeRequest
}

// NewAlibabaAlihealthDoctorLeshuiApplyNotifyRequest 初始化AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest对象
func NewAlibabaAlihealthDoctorLeshuiApplyNotifyRequest() *AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest {
	return &AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.doctor.leshui.apply.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest) SetParam(_param *ApplyTaxNoticeRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest) GetParam() *ApplyTaxNoticeRequest {
	return r._param
}
