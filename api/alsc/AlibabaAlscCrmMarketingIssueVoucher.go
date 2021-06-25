package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
发券 
alibaba.alsc.crm.marketing.issue.voucher

提供发券功能
*/
func AlibabaAlscCrmMarketingIssueVoucher(clt *core.SDKClient, req *alsc.AlibabaAlscCrmMarketingIssueVoucherRequest, session string) (*alsc.AlibabaAlscCrmMarketingIssueVoucherResponse, error) {
    var resp alsc.AlibabaAlscCrmMarketingIssueVoucherAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
