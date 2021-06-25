package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断当前用户是否能对订单执行一些逆向操作，比如退货操作 APIRequest
cainiao.refund.refundactions.judgement

判断当前用户是否能对订单执行一些逆向操作，比如退货操作
*/
type CainiaoRefundRefundactionsJudgementRequest struct {
    model.Params

    // 操作请求
    param0   *OrderRefundOperationJudgementReq 

}

func NewCainiaoRefundRefundactionsJudgementRequest() *CainiaoRefundRefundactionsJudgementRequest{
    return &CainiaoRefundRefundactionsJudgementRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoRefundRefundactionsJudgementRequest) GetApiMethodName() string {
    return "cainiao.refund.refundactions.judgement"
}

func (r CainiaoRefundRefundactionsJudgementRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoRefundRefundactionsJudgementRequest) SetParam0(param0 *OrderRefundOperationJudgementReq) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r CainiaoRefundRefundactionsJudgementRequest) GetParam0() *OrderRefundOperationJudgementReq {
    return r.param0
}

