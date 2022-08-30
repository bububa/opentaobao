package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest 乐税审核结果通知 API请求
// alibaba.alihealth.doctor.leshui.audit.result
//
// 乐税审核结果通知
type AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest struct {
	model.Params
	// 入参
	_param *PayTaxNoticeRequest
}

// NewAlibabaAlihealthDoctorLeshuiAuditResultRequest 初始化AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest对象
func NewAlibabaAlihealthDoctorLeshuiAuditResultRequest() *AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest {
	return &AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.doctor.leshui.audit.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) SetParam(_param *PayTaxNoticeRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest) GetParam() *PayTaxNoticeRequest {
	return r._param
}
