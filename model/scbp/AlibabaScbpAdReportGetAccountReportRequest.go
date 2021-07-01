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
type AlibabaScbpAdReportGetAccountReportAPIRequest struct {
    model.Params
    // 请求参数
    _accountReportOperation   *AccountReportOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdReportGetAccountReportAPIRequest对象
func NewAlibabaScbpAdReportGetAccountReportRequest() *AlibabaScbpAdReportGetAccountReportAPIRequest{
    return &AlibabaScbpAdReportGetAccountReportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetAccountReportAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.get.account.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetAccountReportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportGetAccountReportAPIRequest) SetAccountReportOperation(_accountReportOperation *AccountReportOperationDTO) error {
    r._accountReportOperation = _accountReportOperation
    r.Set("account_report_operation", _accountReportOperation)
    return nil
}

// AccountReportOperation Getter
func (r AlibabaScbpAdReportGetAccountReportAPIRequest) GetAccountReportOperation() *AccountReportOperationDTO {
    return r._accountReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetAccountReportAPIRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportGetAccountReportAPIRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
