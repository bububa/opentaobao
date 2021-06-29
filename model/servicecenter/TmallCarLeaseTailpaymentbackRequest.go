package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
尾款处置方案回传 API请求
tmall.car.lease.tailpaymentback

尾款处置方案回传
*/
type TmallCarLeaseTailpaymentbackRequest struct {
    model.Params
    // 尾款方案
    tailPaymentDTO   *TailPaymentDto
}

// 初始化TmallCarLeaseTailpaymentbackRequest对象
func NewTmallCarLeaseTailpaymentbackRequest() *TmallCarLeaseTailpaymentbackRequest{
    return &TmallCarLeaseTailpaymentbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseTailpaymentbackRequest) GetApiMethodName() string {
    return "tmall.car.lease.tailpaymentback"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseTailpaymentbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TailPaymentDTO Setter
// 尾款方案
func (r *TmallCarLeaseTailpaymentbackRequest) SetTailPaymentDTO(tailPaymentDTO *TailPaymentDto) error {
    r.tailPaymentDTO = tailPaymentDTO
    r.Set("tail_payment_d_t_o", tailPaymentDTO)
    return nil
}

// TailPaymentDTO Getter
func (r TmallCarLeaseTailpaymentbackRequest) GetTailPaymentDTO() *TailPaymentDto {
    return r.tailPaymentDTO
}
