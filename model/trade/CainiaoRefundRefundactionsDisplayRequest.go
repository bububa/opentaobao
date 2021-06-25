package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退货退款操作的展示信息(展现给买家) APIRequest
cainiao.refund.refundactions.display

退货退款操作的展示信息(展现给买家)
*/
type CainiaoRefundRefundactionsDisplayRequest struct {
    model.Params

    // 请求入参
    param0   *OrderRefundOperationReq 

}

func NewCainiaoRefundRefundactionsDisplayRequest() *CainiaoRefundRefundactionsDisplayRequest{
    return &CainiaoRefundRefundactionsDisplayRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoRefundRefundactionsDisplayRequest) GetApiMethodName() string {
    return "cainiao.refund.refundactions.display"
}

func (r CainiaoRefundRefundactionsDisplayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoRefundRefundactionsDisplayRequest) SetParam0(param0 *OrderRefundOperationReq) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r CainiaoRefundRefundactionsDisplayRequest) GetParam0() *OrderRefundOperationReq {
    return r.param0
}

