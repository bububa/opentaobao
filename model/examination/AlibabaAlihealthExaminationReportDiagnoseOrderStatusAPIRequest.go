package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest 报告解读订单状态更新 API请求
// alibaba.alihealth.examination.report.diagnose.order.status
//
// 报告解读订单状态更新
type AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest struct {
	model.Params
	// 参数对象
	_reportOrderStatusRequest *ReportOrderStatusRequest
}

// NewAlibabaalihealthexaminationreportdiagnoseorderstatusRequest 初始化AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest对象
func NewAlibabaalihealthexaminationreportdiagnoseorderstatusRequest() *AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest {
	return &AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportOrderStatusRequest is ReportOrderStatusRequest Setter
// 参数对象
func (r *AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest) SetReportOrderStatusRequest(_reportOrderStatusRequest *ReportOrderStatusRequest) error {
	r._reportOrderStatusRequest = _reportOrderStatusRequest
	r.Set("report_order_status_request", _reportOrderStatusRequest)
	return nil
}

// GetReportOrderStatusRequest ReportOrderStatusRequest Getter
func (r AlibabaalihealthexaminationreportdiagnoseorderstatusAPIRequest) GetReportOrderStatusRequest() *ReportOrderStatusRequest {
	return r._reportOrderStatusRequest
}
