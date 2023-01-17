package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest 获取报告解读url API请求
// alibaba.alihealth.examination.report.diagnose.order.diagnoseurl.get
//
// 获取报告解读url
type AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest struct {
	model.Params
	// 入参
	_param *IsvGetReportDiagnoseUrlRequest
}

// NewAlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetRequest 初始化AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest {
	return &AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.diagnoseurl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest) SetParam(_param *IsvGetReportDiagnoseUrlRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest) GetParam() *IsvGetReportDiagnoseUrlRequest {
	return r._param
}
