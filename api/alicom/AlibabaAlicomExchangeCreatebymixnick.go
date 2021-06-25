package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
代理商积分兑换接口tae 
alibaba.alicom.exchange.createbymixnick

代理商调用该接口来进行积分兑换，tae
*/
func AlibabaAlicomExchangeCreatebymixnick(clt *core.SDKClient, req *alicom.AlibabaAlicomExchangeCreatebymixnickRequest, session string) (*alicom.AlibabaAlicomExchangeCreatebymixnickResponse, error) {
    var resp alicom.AlibabaAlicomExchangeCreatebymixnickAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
