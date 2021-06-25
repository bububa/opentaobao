package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
尾款处置方案回传 APIRequest
tmall.car.lease.tailpaymentback

尾款处置方案回传
*/
type TmallCarLeaseTailpaymentbackRequest struct {
    model.Params

    // 尾款方案
    tailPaymentDTO   *TailPaymentDto 

}

func NewTmallCarLeaseTailpaymentbackRequest() *TmallCarLeaseTailpaymentbackRequest{
    return &TmallCarLeaseTailpaymentbackRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeaseTailpaymentbackRequest) GetApiMethodName() string {
    return "tmall.car.lease.tailpaymentback"
}

func (r TmallCarLeaseTailpaymentbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarLeaseTailpaymentbackRequest) SetTailPaymentDTO(tailPaymentDTO *TailPaymentDto) error {
    r.tailPaymentDTO = tailPaymentDTO
    r.Set("tail_payment_d_t_o", tailPaymentDTO)
    return nil
}

func (r TmallCarLeaseTailpaymentbackRequest) GetTailPaymentDTO() *TailPaymentDto {
    return r.tailPaymentDTO
}

