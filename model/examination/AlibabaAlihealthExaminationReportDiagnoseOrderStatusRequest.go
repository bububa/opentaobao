package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
报告解读订单状态更新 API请求
alibaba.alihealth.examination.report.diagnose.order.status

报告解读订单状态更新
*/
type AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest struct {
    model.Params
    // 参数对象
    _reportOrderStatusRequest   *ReportOrderStatusRequest
}

// 初始化AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest{
    return &AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.report.diagnose.order.status"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReportOrderStatusRequest Setter
// 参数对象
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest) SetReportOrderStatusRequest(_reportOrderStatusRequest *ReportOrderStatusRequest) error {
    r._reportOrderStatusRequest = _reportOrderStatusRequest
    r.Set("report_order_status_request", _reportOrderStatusRequest)
    return nil
}

// ReportOrderStatusRequest Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest) GetReportOrderStatusRequest() *ReportOrderStatusRequest {
    return r._reportOrderStatusRequest
}
