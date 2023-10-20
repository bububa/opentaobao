package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest 获取报告解读url API请求
// alibaba.alihealth.examination.report.diagnose.order.diagnoseurl.get
//
// 获取报告解读url
type AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest struct {
	model.Params
	// 入参
	_param *IsvGetReportDiagnoseUrlRequest
}

// NewAlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetRequest 初始化AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest对象
func NewAlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetRequest() *AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest {
	return &AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.diagnoseurl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest) SetParam(_param *IsvGetReportDiagnoseUrlRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaalihealthexaminationreportdiagnoseorderdiagnoseurlgetAPIRequest) GetParam() *IsvGetReportDiagnoseUrlRequest {
	return r._param
}
