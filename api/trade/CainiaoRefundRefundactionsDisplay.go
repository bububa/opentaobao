package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
退货退款操作的展示信息(展现给买家) 
cainiao.refund.refundactions.display

退货退款操作的展示信息(展现给买家)
*/
func CainiaoRefundRefundactionsDisplay(clt *core.SDKClient, req *trade.CainiaoRefundRefundactionsDisplayAPIRequest, session string) (*trade.CainiaoRefundRefundactionsDisplayAPIResponse, error) {
    var resp trade.CainiaoRefundRefundactionsDisplayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
