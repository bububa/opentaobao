package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
推客平台订单回流 
alibaba.trade.aliance.create

推客平台订单回流
*/
func AlibabaTradeAlianceCreate(clt *core.SDKClient, req *trade.AlibabaTradeAlianceCreateRequest, session string) (*trade.AlibabaTradeAlianceCreateResponse, error) {
    var resp trade.AlibabaTradeAlianceCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
