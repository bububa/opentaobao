package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定报告 APIRequest
alibaba.idle.recycle.inspection.report

回收商鉴定报告
*/
type AlibabaIdleRecycleInspectionReportRequest struct {
    model.Params

    // 鉴定报告
    inspectionReport   *InspectionReport 

}

func NewAlibabaIdleRecycleInspectionReportRequest() *AlibabaIdleRecycleInspectionReportRequest{
    return &AlibabaIdleRecycleInspectionReportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRecycleInspectionReportRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.inspection.report"
}

func (r AlibabaIdleRecycleInspectionReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRecycleInspectionReportRequest) SetInspectionReport(inspectionReport *InspectionReport) error {
    r.inspectionReport = inspectionReport
    r.Set("inspection_report", inspectionReport)
    return nil
}

func (r AlibabaIdleRecycleInspectionReportRequest) GetInspectionReport() *InspectionReport {
    return r.inspectionReport
}

