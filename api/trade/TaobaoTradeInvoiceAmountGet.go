package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
获取订单应开票金额 
taobao.trade.invoice.amount.get

订单应开票金额计算
*/
func TaobaoTradeInvoiceAmountGet(clt *core.SDKClient, req *trade.TaobaoTradeInvoiceAmountGetRequest, session string) (*trade.TaobaoTradeInvoiceAmountGetResponse, error) {
    var resp trade.TaobaoTradeInvoiceAmountGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
