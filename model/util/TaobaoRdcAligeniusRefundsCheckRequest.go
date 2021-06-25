package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款信息审核 APIRequest
taobao.rdc.aligenius.refunds.check

根据退款信息，对退款单进行审核
*/
type TaobaoRdcAligeniusRefundsCheckRequest struct {
    model.Params

    // 入参
    param   *RefundCheckDto 

}

func NewTaobaoRdcAligeniusRefundsCheckRequest() *TaobaoRdcAligeniusRefundsCheckRequest{
    return &TaobaoRdcAligeniusRefundsCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdcAligeniusRefundsCheckRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.refunds.check"
}

func (r TaobaoRdcAligeniusRefundsCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdcAligeniusRefundsCheckRequest) SetParam(param *RefundCheckDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoRdcAligeniusRefundsCheckRequest) GetParam() *RefundCheckDto {
    return r.param
}

