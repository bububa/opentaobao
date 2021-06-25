package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消付款单 APIRequest
alibaba.mj.mos.fund.cancelbill

取消付款单
*/
type AlibabaMjMosFundCancelbillRequest struct {
    model.Params

    // 取消入参
    cancelBillDTO   *CancelBillDto 

}

func NewAlibabaMjMosFundCancelbillRequest() *AlibabaMjMosFundCancelbillRequest{
    return &AlibabaMjMosFundCancelbillRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjMosFundCancelbillRequest) GetApiMethodName() string {
    return "alibaba.mj.mos.fund.cancelbill"
}

func (r AlibabaMjMosFundCancelbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjMosFundCancelbillRequest) SetCancelBillDTO(cancelBillDTO *CancelBillDto) error {
    r.cancelBillDTO = cancelBillDTO
    r.Set("cancel_bill_d_t_o", cancelBillDTO)
    return nil
}

func (r AlibabaMjMosFundCancelbillRequest) GetCancelBillDTO() *CancelBillDto {
    return r.cancelBillDTO
}

