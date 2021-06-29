package lsttrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lsttrade"
)

/* 
零售通交易订单查询--品牌商视角 
alibaba.lst.trade.order.get

根据订单id查询零售通交易订单
*/
func AlibabaLstTradeOrderGet(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeOrderGetRequest, session string) (*lsttrade.AlibabaLstTradeOrderGetAPIResponse, error) {
    var resp lsttrade.AlibabaLstTradeOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}