package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改付款单的银行账户信息 APIRequest
alibaba.mj.mos.fund.modifybillbankaccount

修改付款单的银行账户信息
*/
type AlibabaMjMosFundModifybillbankaccountRequest struct {
    model.Params

    // 修改入参
    modifyDto   *ModifyBillDto 

}

func NewAlibabaMjMosFundModifybillbankaccountRequest() *AlibabaMjMosFundModifybillbankaccountRequest{
    return &AlibabaMjMosFundModifybillbankaccountRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjMosFundModifybillbankaccountRequest) GetApiMethodName() string {
    return "alibaba.mj.mos.fund.modifybillbankaccount"
}

func (r AlibabaMjMosFundModifybillbankaccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjMosFundModifybillbankaccountRequest) SetModifyDto(modifyDto *ModifyBillDto) error {
    r.modifyDto = modifyDto
    r.Set("modify_dto", modifyDto)
    return nil
}

func (r AlibabaMjMosFundModifybillbankaccountRequest) GetModifyDto() *ModifyBillDto {
    return r.modifyDto
}

