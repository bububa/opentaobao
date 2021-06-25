package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向报告 APIRequest
alibaba.scbp.ad.report.get.target.report

定向报告
*/
type AlibabaScbpAdReportGetTargetReportRequest struct {
    model.Params

    // 请求参数
    targetReportOperation   *TargetReportOperationDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdReportGetTargetReportRequest() *AlibabaScbpAdReportGetTargetReportRequest{
    return &AlibabaScbpAdReportGetTargetReportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdReportGetTargetReportRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.get.target.report"
}

func (r AlibabaScbpAdReportGetTargetReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdReportGetTargetReportRequest) SetTargetReportOperation(targetReportOperation *TargetReportOperationDto) error {
    r.targetReportOperation = targetReportOperation
    r.Set("target_report_operation", targetReportOperation)
    return nil
}

func (r AlibabaScbpAdReportGetTargetReportRequest) GetTargetReportOperation() *TargetReportOperationDto {
    return r.targetReportOperation
}

func (r *AlibabaScbpAdReportGetTargetReportRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdReportGetTargetReportRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

