package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询店仓作业单据清单 （库存对账辅助）-回流单 APIRequest
alibaba.wdk.ums.order.get

查询店仓作业单据清单 （库存对账辅助）-回流单
*/
type AlibabaWdkUmsOrderGetRequest struct {
    model.Params

    // 查询单据的dto
    queryErpbillDto   *QueryErpBillDto 

}

func NewAlibabaWdkUmsOrderGetRequest() *AlibabaWdkUmsOrderGetRequest{
    return &AlibabaWdkUmsOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsOrderGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.order.get"
}

func (r AlibabaWdkUmsOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsOrderGetRequest) SetQueryErpbillDto(queryErpbillDto *QueryErpBillDto) error {
    r.queryErpbillDto = queryErpbillDto
    r.Set("query_erpbill_dto", queryErpbillDto)
    return nil
}

func (r AlibabaWdkUmsOrderGetRequest) GetQueryErpbillDto() *QueryErpBillDto {
    return r.queryErpbillDto
}

