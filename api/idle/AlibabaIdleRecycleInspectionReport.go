package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
鉴定报告 
alibaba.idle.recycle.inspection.report

回收商鉴定报告
*/
func AlibabaIdleRecycleInspectionReport(clt *core.SDKClient, req *idle.AlibabaIdleRecycleInspectionReportRequest, session string) (*idle.AlibabaIdleRecycleInspectionReportAPIResponse, error) {
    var resp idle.AlibabaIdleRecycleInspectionReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
