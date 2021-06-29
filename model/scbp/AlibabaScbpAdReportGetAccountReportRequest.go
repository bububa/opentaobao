package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
账户报告 API请求
alibaba.scbp.ad.report.get.account.report

账户报告
*/
type AlibabaScbpAdReportGetAccountReportRequest struct {
    model.Params
    // 请求参数
    accountReportOperation   *AccountReportOperationDto
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdReportGetAccountReportRequest对象
func NewAlibabaScbpAdReportGetAccountReportRequest() *AlibabaScbpAdReportGetAccountReportRequest{
    return &AlibabaScbpAdReportGetAccountReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetAccountReportRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.get.account.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetAccountReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportGetAccountReportRequest) SetAccountReportOperation(accountReportOperation *AccountReportOperationDto) error {
    r.accountReportOperation = accountReportOperation
    r.Set("account_report_operation", accountReportOperation)
    return nil
}

// AccountReportOperation Getter
func (r AlibabaScbpAdReportGetAccountReportRequest) GetAccountReportOperation() *AccountReportOperationDto {
    return r.accountReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetAccountReportRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportGetAccountReportRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
