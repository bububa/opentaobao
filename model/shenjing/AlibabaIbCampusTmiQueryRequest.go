package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IB智慧园区-查询TMI流水 APIRequest
alibaba.ib.campus.tmi.query

获取特定银行账户的银行流水
*/
type AlibabaIbCampusTmiQueryRequest struct {
    model.Params

    // 查询参数
    accountQueryReqDto   *AccountQueryReqDto 

}

func NewAlibabaIbCampusTmiQueryRequest() *AlibabaIbCampusTmiQueryRequest{
    return &AlibabaIbCampusTmiQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIbCampusTmiQueryRequest) GetApiMethodName() string {
    return "alibaba.ib.campus.tmi.query"
}

func (r AlibabaIbCampusTmiQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIbCampusTmiQueryRequest) SetAccountQueryReqDto(accountQueryReqDto *AccountQueryReqDto) error {
    r.accountQueryReqDto = accountQueryReqDto
    r.Set("account_query_req_dto", accountQueryReqDto)
    return nil
}

func (r AlibabaIbCampusTmiQueryRequest) GetAccountQueryReqDto() *AccountQueryReqDto {
    return r.accountQueryReqDto
}

