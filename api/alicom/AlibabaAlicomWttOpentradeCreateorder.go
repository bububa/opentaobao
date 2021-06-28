package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
充值送活动下单接口 
alibaba.alicom.wtt.opentrade.createorder

提供给话费宝创建淘宝订单
*/
func AlibabaAlicomWttOpentradeCreateorder(clt *core.SDKClient, req *alicom.AlibabaAlicomWttOpentradeCreateorderRequest, session string) (*alicom.AlibabaAlicomWttOpentradeCreateorderAPIResponse, error) {
    var resp alicom.AlibabaAlicomWttOpentradeCreateorderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
