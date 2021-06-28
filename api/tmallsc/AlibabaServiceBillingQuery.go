package tmallsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallsc"
)

/* 
服务平台结算出账信息 
alibaba.service.billing.query

服务平台结算单明细查询服务
*/
func AlibabaServiceBillingQuery(clt *core.SDKClient, req *tmallsc.AlibabaServiceBillingQueryRequest, session string) (*tmallsc.AlibabaServiceBillingQueryAPIResponse, error) {
    var resp tmallsc.AlibabaServiceBillingQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
