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
    _accountReportOperation   *AccountReportOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
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
func (r *AlibabaScbpAdReportGetAccountReportRequest) SetAccountReportOperation(_accountReportOperation *AccountReportOperationDTO) error {
    r._accountReportOperation = _accountReportOperation
    r.Set("account_report_operation", _accountReportOperation)
    return nil
}

// AccountReportOperation Getter
func (r AlibabaScbpAdReportGetAccountReportRequest) GetAccountReportOperation() *AccountReportOperationDTO {
    return r._accountReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetAccountReportRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportGetAccountReportRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
