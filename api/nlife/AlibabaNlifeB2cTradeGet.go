package nlife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nlife"
)

/* 
零售+平台查询订单 
alibaba.nlife.b2c.trade.get

查询零售+平台创建出来的订单详情
*/
func AlibabaNlifeB2cTradeGet(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradeGetAPIRequest, session string) (*nlife.AlibabaNlifeB2cTradeGetAPIResponse, error) {
    var resp nlife.AlibabaNlifeB2cTradeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
