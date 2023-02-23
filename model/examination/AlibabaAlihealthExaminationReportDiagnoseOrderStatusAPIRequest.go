package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest 报告解读订单状态更新 API请求
// alibaba.alihealth.examination.report.diagnose.order.status
//
// 报告解读订单状态更新
type AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest struct {
	model.Params
	// 参数对象
	_reportOrderStatusRequest *ReportOrderStatusRequest
}

// NewAlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest 初始化AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest {
	return &AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportOrderStatusRequest is ReportOrderStatusRequest Setter
// 参数对象
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest) SetReportOrderStatusRequest(_reportOrderStatusRequest *ReportOrderStatusRequest) error {
	r._reportOrderStatusRequest = _reportOrderStatusRequest
	r.Set("report_order_status_request", _reportOrderStatusRequest)
	return nil
}

// GetReportOrderStatusRequest ReportOrderStatusRequest Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest) GetReportOrderStatusRequest() *ReportOrderStatusRequest {
	return r._reportOrderStatusRequest
}
