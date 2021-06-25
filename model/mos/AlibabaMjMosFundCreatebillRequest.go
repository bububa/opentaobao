package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一个付款单 APIRequest
alibaba.mj.mos.fund.createbill

创建一个付款单
*/
type AlibabaMjMosFundCreatebillRequest struct {
    model.Params

    // 创建付款单入参
    billDto   *CreateBillDto 

}

func NewAlibabaMjMosFundCreatebillRequest() *AlibabaMjMosFundCreatebillRequest{
    return &AlibabaMjMosFundCreatebillRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjMosFundCreatebillRequest) GetApiMethodName() string {
    return "alibaba.mj.mos.fund.createbill"
}

func (r AlibabaMjMosFundCreatebillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjMosFundCreatebillRequest) SetBillDto(billDto *CreateBillDto) error {
    r.billDto = billDto
    r.Set("bill_dto", billDto)
    return nil
}

func (r AlibabaMjMosFundCreatebillRequest) GetBillDto() *CreateBillDto {
    return r.billDto
}

