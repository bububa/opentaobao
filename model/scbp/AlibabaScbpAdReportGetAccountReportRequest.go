package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
账户报告 APIRequest
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

func NewAlibabaScbpAdReportGetAccountReportRequest() *AlibabaScbpAdReportGetAccountReportRequest{
    return &AlibabaScbpAdReportGetAccountReportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdReportGetAccountReportRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.get.account.report"
}

func (r AlibabaScbpAdReportGetAccountReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdReportGetAccountReportRequest) SetAccountReportOperation(accountReportOperation *AccountReportOperationDto) error {
    r.accountReportOperation = accountReportOperation
    r.Set("account_report_operation", accountReportOperation)
    return nil
}

func (r AlibabaScbpAdReportGetAccountReportRequest) GetAccountReportOperation() *AccountReportOperationDto {
    return r.accountReportOperation
}

func (r *AlibabaScbpAdReportGetAccountReportRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdReportGetAccountReportRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

