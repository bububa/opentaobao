package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
确认收款 
taobao.fenxiao.order.confirm.paid

供应商确认收款（非支付宝交易）。
*/
func TaobaoFenxiaoOrderConfirmPaid(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoOrderConfirmPaidAPIRequest, session string) (*fenxiao.TaobaoFenxiaoOrderConfirmPaidAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoOrderConfirmPaidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
