package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest 报告解读令牌校验 API请求
// alibaba.alihealth.examination.report.diagnose.order.verify
//
// 报告解读令牌校验
type AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest struct {
	model.Params
	// 请求对象
	_reportImTokenStatusRequest *ReportImTokenStatusRequest
}

// NewAlibabaAlihealthExaminationReportDiagnoseOrderVerifyRequest 初始化AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseOrderVerifyRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest {
	return &AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportImTokenStatusRequest is ReportImTokenStatusRequest Setter
// 请求对象
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) SetReportImTokenStatusRequest(_reportImTokenStatusRequest *ReportImTokenStatusRequest) error {
	r._reportImTokenStatusRequest = _reportImTokenStatusRequest
	r.Set("report_im_token_status_request", _reportImTokenStatusRequest)
	return nil
}

// GetReportImTokenStatusRequest ReportImTokenStatusRequest Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) GetReportImTokenStatusRequest() *ReportImTokenStatusRequest {
	return r._reportImTokenStatusRequest
}
