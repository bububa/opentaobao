package lstmarketing

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstmarketing"
)

/* 
根据订单查询营销信息 
alibaba.lst.marketing.querybyorderid

根据订单查询营销信息
*/
func AlibabaLstMarketingQuerybyorderid(clt *core.SDKClient, req *lstmarketing.AlibabaLstMarketingQuerybyorderidRequest, session string) (*lstmarketing.AlibabaLstMarketingQuerybyorderidAPIResponse, error) {
    var resp lstmarketing.AlibabaLstMarketingQuerybyorderidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
