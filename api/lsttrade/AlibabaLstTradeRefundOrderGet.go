package lsttrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lsttrade"
)

/* 
零售通退款订单查询 
alibaba.lst.trade.refund.order.get

零售通退款订单查询
*/
func AlibabaLstTradeRefundOrderGet(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeRefundOrderGetRequest, session string) (*lsttrade.AlibabaLstTradeRefundOrderGetAPIResponse, error) {
    var resp lsttrade.AlibabaLstTradeRefundOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
