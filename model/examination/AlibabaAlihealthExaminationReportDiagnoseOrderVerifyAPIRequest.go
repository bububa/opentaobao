package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest 报告解读令牌校验 API请求
// alibaba.alihealth.examination.report.diagnose.order.verify
//
// 报告解读令牌校验
type AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest struct {
	model.Params
	// 请求对象
	_reportImTokenStatusRequest *ReportImTokenStatusRequest
}

// NewAlibabaalihealthexaminationreportdiagnoseorderverifyRequest 初始化AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest对象
func NewAlibabaalihealthexaminationreportdiagnoseorderverifyRequest() *AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest {
	return &AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportImTokenStatusRequest is ReportImTokenStatusRequest Setter
// 请求对象
func (r *AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest) SetReportImTokenStatusRequest(_reportImTokenStatusRequest *ReportImTokenStatusRequest) error {
	r._reportImTokenStatusRequest = _reportImTokenStatusRequest
	r.Set("report_im_token_status_request", _reportImTokenStatusRequest)
	return nil
}

// GetReportImTokenStatusRequest ReportImTokenStatusRequest Getter
func (r AlibabaalihealthexaminationreportdiagnoseorderverifyAPIRequest) GetReportImTokenStatusRequest() *ReportImTokenStatusRequest {
	return r._reportImTokenStatusRequest
}
