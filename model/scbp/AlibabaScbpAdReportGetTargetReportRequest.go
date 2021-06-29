package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向报告 API请求
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

// 初始化AlibabaScbpAdReportGetTargetReportRequest对象
func NewAlibabaScbpAdReportGetTargetReportRequest() *AlibabaScbpAdReportGetTargetReportRequest{
    return &AlibabaScbpAdReportGetTargetReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetTargetReportRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.get.target.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetTargetReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TargetReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportGetTargetReportRequest) SetTargetReportOperation(targetReportOperation *TargetReportOperationDto) error {
    r.targetReportOperation = targetReportOperation
    r.Set("target_report_operation", targetReportOperation)
    return nil
}

// TargetReportOperation Getter
func (r AlibabaScbpAdReportGetTargetReportRequest) GetTargetReportOperation() *TargetReportOperationDto {
    return r.targetReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetTargetReportRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportGetTargetReportRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
