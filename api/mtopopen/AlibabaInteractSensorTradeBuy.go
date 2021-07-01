package mtopopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mtopopen"
)

/* 
手淘下单能力开放 
alibaba.interact.sensor.trade.buy

交易流程鉴权
*/
func AlibabaInteractSensorTradeBuy(clt *core.SDKClient, req *mtopopen.AlibabaInteractSensorTradeBuyAPIRequest, session string) (*mtopopen.AlibabaInteractSensorTradeBuyAPIResponse, error) {
    var resp mtopopen.AlibabaInteractSensorTradeBuyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
