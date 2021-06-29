package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
报告解读订单状态更新 APIRequest
alibaba.alihealth.examination.report.diagnose.order.status

报告解读订单状态更新
*/
type AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest struct {
    model.Params

    // 参数对象
    reportOrderStatusRequest   *ReportOrderStatusRequest 

}

func NewAlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest{
    return &AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.report.diagnose.order.status"
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest) SetReportOrderStatusRequest(reportOrderStatusRequest *ReportOrderStatusRequest) error {
    r.reportOrderStatusRequest = reportOrderStatusRequest
    r.Set("report_order_status_request", reportOrderStatusRequest)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderStatusRequest) GetReportOrderStatusRequest() *ReportOrderStatusRequest {
    return r.reportOrderStatusRequest
}

