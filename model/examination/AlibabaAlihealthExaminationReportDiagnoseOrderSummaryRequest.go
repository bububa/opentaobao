package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检报告人工解读总结回传 API请求
alibaba.alihealth.examination.report.diagnose.order.summary

记录体检报告人工解读总结
*/
type AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest struct {
    model.Params
    // 入参对象
    reportOrderSummaryRequest   *ReportOrderSummaryRequest
}

// 初始化AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest{
    return &AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.report.diagnose.order.summary"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReportOrderSummaryRequest Setter
// 入参对象
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest) SetReportOrderSummaryRequest(reportOrderSummaryRequest *ReportOrderSummaryRequest) error {
    r.reportOrderSummaryRequest = reportOrderSummaryRequest
    r.Set("report_order_summary_request", reportOrderSummaryRequest)
    return nil
}

// ReportOrderSummaryRequest Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest) GetReportOrderSummaryRequest() *ReportOrderSummaryRequest {
    return r.reportOrderSummaryRequest
}
