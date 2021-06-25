package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
积分兑换校验淘宝账号是否存在 
alibaba.wt.order.exchange.partner.checktbuser

积分兑换之前校验淘宝账号是否存在
*/
func AlibabaWtOrderExchangePartnerChecktbuser(clt *core.SDKClient, req *alicom.AlibabaWtOrderExchangePartnerChecktbuserRequest, session string) (*alicom.AlibabaWtOrderExchangePartnerChecktbuserResponse, error) {
    var resp alicom.AlibabaWtOrderExchangePartnerChecktbuserAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
