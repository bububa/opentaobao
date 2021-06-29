package lsttrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lsttrade"
)

/* 
查询退款单列表(卖家视角) 
alibaba.lst.trade.order.refund.list.query

查询退款单列表(卖家视角)
*/
func AlibabaLstTradeOrderRefundListQuery(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeOrderRefundListQueryRequest, session string) (*lsttrade.AlibabaLstTradeOrderRefundListQueryAPIResponse, error) {
    var resp lsttrade.AlibabaLstTradeOrderRefundListQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
