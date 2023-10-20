package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest 体检报告人工解读总结回传 API请求
// alibaba.alihealth.examination.report.diagnose.order.summary
//
// 记录体检报告人工解读总结
type AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest struct {
	model.Params
	// 入参对象
	_reportOrderSummaryRequest *ReportOrderSummaryRequest
}

// NewAlibabaalihealthexaminationreportdiagnoseordersummaryRequest 初始化AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest对象
func NewAlibabaalihealthexaminationreportdiagnoseordersummaryRequest() *AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest {
	return &AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.summary"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportOrderSummaryRequest is ReportOrderSummaryRequest Setter
// 入参对象
func (r *AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest) SetReportOrderSummaryRequest(_reportOrderSummaryRequest *ReportOrderSummaryRequest) error {
	r._reportOrderSummaryRequest = _reportOrderSummaryRequest
	r.Set("report_order_summary_request", _reportOrderSummaryRequest)
	return nil
}

// GetReportOrderSummaryRequest ReportOrderSummaryRequest Getter
func (r AlibabaalihealthexaminationreportdiagnoseordersummaryAPIRequest) GetReportOrderSummaryRequest() *ReportOrderSummaryRequest {
	return r._reportOrderSummaryRequest
}
