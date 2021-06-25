package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
定向报告 
alibaba.scbp.ad.report.get.target.report

定向报告
*/
func AlibabaScbpAdReportGetTargetReport(clt *core.SDKClient, req *scbp.AlibabaScbpAdReportGetTargetReportRequest, session string) (*scbp.AlibabaScbpAdReportGetTargetReportResponse, error) {
    var resp scbp.AlibabaScbpAdReportGetTargetReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
