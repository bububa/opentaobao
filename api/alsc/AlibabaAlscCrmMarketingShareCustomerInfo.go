package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
查询分享营销客户领券信息 
alibaba.alsc.crm.marketing.share.customer.info

查询分享营销活动的客户领券信息
*/
func AlibabaAlscCrmMarketingShareCustomerInfo(clt *core.SDKClient, req *alsc.AlibabaAlscCrmMarketingShareCustomerInfoRequest, session string) (*alsc.AlibabaAlscCrmMarketingShareCustomerInfoResponse, error) {
    var resp alsc.AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
