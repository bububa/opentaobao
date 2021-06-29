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
    _targetReportOperation   *TargetReportOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
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
func (r *AlibabaScbpAdReportGetTargetReportRequest) SetTargetReportOperation(_targetReportOperation *TargetReportOperationDTO) error {
    r._targetReportOperation = _targetReportOperation
    r.Set("target_report_operation", _targetReportOperation)
    return nil
}

// TargetReportOperation Getter
func (r AlibabaScbpAdReportGetTargetReportRequest) GetTargetReportOperation() *TargetReportOperationDTO {
    return r._targetReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetTargetReportRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportGetTargetReportRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
