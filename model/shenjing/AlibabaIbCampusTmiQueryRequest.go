package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IB智慧园区-查询TMI流水 API请求
alibaba.ib.campus.tmi.query

获取特定银行账户的银行流水
*/
type AlibabaIbCampusTmiQueryRequest struct {
    model.Params
    // 查询参数
    accountQueryReqDto   *AccountQueryReqDto
}

// 初始化AlibabaIbCampusTmiQueryRequest对象
func NewAlibabaIbCampusTmiQueryRequest() *AlibabaIbCampusTmiQueryRequest{
    return &AlibabaIbCampusTmiQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIbCampusTmiQueryRequest) GetApiMethodName() string {
    return "alibaba.ib.campus.tmi.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIbCampusTmiQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountQueryReqDto Setter
// 查询参数
func (r *AlibabaIbCampusTmiQueryRequest) SetAccountQueryReqDto(accountQueryReqDto *AccountQueryReqDto) error {
    r.accountQueryReqDto = accountQueryReqDto
    r.Set("account_query_req_dto", accountQueryReqDto)
    return nil
}

// AccountQueryReqDto Getter
func (r AlibabaIbCampusTmiQueryRequest) GetAccountQueryReqDto() *AccountQueryReqDto {
    return r.accountQueryReqDto
}
