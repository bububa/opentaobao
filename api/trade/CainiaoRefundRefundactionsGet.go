package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
判断该订单能执行的逆向操作集合列表 
cainiao.refund.refundactions.get

判断该订单能执行的逆向操作集合列表
*/
func CainiaoRefundRefundactionsGet(clt *core.SDKClient, req *trade.CainiaoRefundRefundactionsGetRequest, session string) (*trade.CainiaoRefundRefundactionsGetAPIResponse, error) {
    var resp trade.CainiaoRefundRefundactionsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
