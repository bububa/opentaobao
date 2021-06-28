package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
退票接口 
alitrip.rail.trade.refund

退票接口
*/
func AlitripRailTradeRefund(clt *core.SDKClient, req *trade.AlitripRailTradeRefundRequest, session string) (*trade.AlitripRailTradeRefundAPIResponse, error) {
    var resp trade.AlitripRailTradeRefundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
