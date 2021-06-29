package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检报告人工解读总结回传 APIRequest
alibaba.alihealth.examination.report.diagnose.order.summary

记录体检报告人工解读总结
*/
type AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest struct {
    model.Params

    // 入参对象
    reportOrderSummaryRequest   *ReportOrderSummaryRequest 

}

func NewAlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest{
    return &AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.report.diagnose.order.summary"
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest) SetReportOrderSummaryRequest(reportOrderSummaryRequest *ReportOrderSummaryRequest) error {
    r.reportOrderSummaryRequest = reportOrderSummaryRequest
    r.Set("report_order_summary_request", reportOrderSummaryRequest)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest) GetReportOrderSummaryRequest() *ReportOrderSummaryRequest {
    return r.reportOrderSummaryRequest
}

