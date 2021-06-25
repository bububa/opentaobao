package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
判断当前用户是否能对订单执行一些逆向操作，比如退货操作 
cainiao.refund.refundactions.judgement

判断当前用户是否能对订单执行一些逆向操作，比如退货操作
*/
func CainiaoRefundRefundactionsJudgement(clt *core.SDKClient, req *trade.CainiaoRefundRefundactionsJudgementRequest, session string) (*trade.CainiaoRefundRefundactionsJudgementResponse, error) {
    var resp trade.CainiaoRefundRefundactionsJudgementAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
