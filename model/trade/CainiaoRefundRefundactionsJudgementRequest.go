package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断当前用户是否能对订单执行一些逆向操作，比如退货操作 API请求
cainiao.refund.refundactions.judgement

判断当前用户是否能对订单执行一些逆向操作，比如退货操作
*/
type CainiaoRefundRefundactionsJudgementRequest struct {
    model.Params
    // 操作请求
    _param0   *OrderRefundOperationJudgementReq
}

// 初始化CainiaoRefundRefundactionsJudgementRequest对象
func NewCainiaoRefundRefundactionsJudgementRequest() *CainiaoRefundRefundactionsJudgementRequest{
    return &CainiaoRefundRefundactionsJudgementRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoRefundRefundactionsJudgementRequest) GetApiMethodName() string {
    return "cainiao.refund.refundactions.judgement"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoRefundRefundactionsJudgementRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 操作请求
func (r *CainiaoRefundRefundactionsJudgementRequest) SetParam0(_param0 *OrderRefundOperationJudgementReq) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r CainiaoRefundRefundactionsJudgementRequest) GetParam0() *OrderRefundOperationJudgementReq {
    return r._param0
}
