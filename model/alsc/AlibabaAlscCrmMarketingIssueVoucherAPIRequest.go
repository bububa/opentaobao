package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发券 API请求
alibaba.alsc.crm.marketing.issue.voucher

提供发券功能
*/
type AlibabaAlscCrmMarketingIssueVoucherAPIRequest struct {
    model.Params
    // 参数
    _paramIssueVoucherReq   *IssueVoucherReq
}

// 初始化AlibabaAlscCrmMarketingIssueVoucherAPIRequest对象
func NewAlibabaAlscCrmMarketingIssueVoucherRequest() *AlibabaAlscCrmMarketingIssueVoucherAPIRequest{
    return &AlibabaAlscCrmMarketingIssueVoucherAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmMarketingIssueVoucherAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.marketing.issue.voucher"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmMarketingIssueVoucherAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamIssueVoucherReq Setter
// 参数
func (r *AlibabaAlscCrmMarketingIssueVoucherAPIRequest) SetParamIssueVoucherReq(_paramIssueVoucherReq *IssueVoucherReq) error {
    r._paramIssueVoucherReq = _paramIssueVoucherReq
    r.Set("param_issue_voucher_req", _paramIssueVoucherReq)
    return nil
}

// ParamIssueVoucherReq Getter
func (r AlibabaAlscCrmMarketingIssueVoucherAPIRequest) GetParamIssueVoucherReq() *IssueVoucherReq {
    return r._paramIssueVoucherReq
}
