package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定报告 API请求
alibaba.idle.recycle.inspection.report

回收商鉴定报告
*/
type AlibabaIdleRecycleInspectionReportRequest struct {
    model.Params
    // 鉴定报告
    inspectionReport   *InspectionReport
}

// 初始化AlibabaIdleRecycleInspectionReportRequest对象
func NewAlibabaIdleRecycleInspectionReportRequest() *AlibabaIdleRecycleInspectionReportRequest{
    return &AlibabaIdleRecycleInspectionReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleInspectionReportRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.inspection.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleInspectionReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InspectionReport Setter
// 鉴定报告
func (r *AlibabaIdleRecycleInspectionReportRequest) SetInspectionReport(inspectionReport *InspectionReport) error {
    r.inspectionReport = inspectionReport
    r.Set("inspection_report", inspectionReport)
    return nil
}

// InspectionReport Getter
func (r AlibabaIdleRecycleInspectionReportRequest) GetInspectionReport() *InspectionReport {
    return r.inspectionReport
}
