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
type AlibabaIdleRecycleInspectionReportAPIRequest struct {
    model.Params
    // 鉴定报告
    _inspectionReport   *InspectionReport
}

// 初始化AlibabaIdleRecycleInspectionReportAPIRequest对象
func NewAlibabaIdleRecycleInspectionReportRequest() *AlibabaIdleRecycleInspectionReportAPIRequest{
    return &AlibabaIdleRecycleInspectionReportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleInspectionReportAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.inspection.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleInspectionReportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InspectionReport Setter
// 鉴定报告
func (r *AlibabaIdleRecycleInspectionReportAPIRequest) SetInspectionReport(_inspectionReport *InspectionReport) error {
    r._inspectionReport = _inspectionReport
    r.Set("inspection_report", _inspectionReport)
    return nil
}

// InspectionReport Getter
func (r AlibabaIdleRecycleInspectionReportAPIRequest) GetInspectionReport() *InspectionReport {
    return r._inspectionReport
}
