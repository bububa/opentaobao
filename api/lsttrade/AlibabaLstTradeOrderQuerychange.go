package lsttrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lsttrade"
)

/* 
订单id批量查询（品牌商视角） 
alibaba.lst.trade.order.querychange

根据品牌和时间段查询有变更记录的订单id
*/
func AlibabaLstTradeOrderQuerychange(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeOrderQuerychangeRequest, session string) (*lsttrade.AlibabaLstTradeOrderQuerychangeAPIResponse, error) {
    var resp lsttrade.AlibabaLstTradeOrderQuerychangeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
