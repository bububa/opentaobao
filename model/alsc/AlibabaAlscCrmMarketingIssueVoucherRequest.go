package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发券 APIRequest
alibaba.alsc.crm.marketing.issue.voucher

提供发券功能
*/
type AlibabaAlscCrmMarketingIssueVoucherRequest struct {
    model.Params

    // 参数
    paramIssueVoucherReq   *IssueVoucherReq 

}

func NewAlibabaAlscCrmMarketingIssueVoucherRequest() *AlibabaAlscCrmMarketingIssueVoucherRequest{
    return &AlibabaAlscCrmMarketingIssueVoucherRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmMarketingIssueVoucherRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.marketing.issue.voucher"
}

func (r AlibabaAlscCrmMarketingIssueVoucherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmMarketingIssueVoucherRequest) SetParamIssueVoucherReq(paramIssueVoucherReq *IssueVoucherReq) error {
    r.paramIssueVoucherReq = paramIssueVoucherReq
    r.Set("param_issue_voucher_req", paramIssueVoucherReq)
    return nil
}

func (r AlibabaAlscCrmMarketingIssueVoucherRequest) GetParamIssueVoucherReq() *IssueVoucherReq {
    return r.paramIssueVoucherReq
}

