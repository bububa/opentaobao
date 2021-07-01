package nlife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nlife"
)

/* 
零售+平台取消订单 
alibaba.nlife.b2c.trade.cancel

零售+平台取消订单接口
*/
func AlibabaNlifeB2cTradeCancel(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradeCancelAPIRequest, session string) (*nlife.AlibabaNlifeB2cTradeCancelAPIResponse, error) {
    var resp nlife.AlibabaNlifeB2cTradeCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
