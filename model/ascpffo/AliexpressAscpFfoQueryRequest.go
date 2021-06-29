package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress发货单查询API APIRequest
aliexpress.ascp.ffo.query

AE 履约发货单分页查询接口
*/
type AliexpressAscpFfoQueryRequest struct {
    model.Params

    // dto
    fulfillmentForwardOrderQuery   *FulfillmentForwardOrderQueryDto 

}

func NewAliexpressAscpFfoQueryRequest() *AliexpressAscpFfoQueryRequest{
    return &AliexpressAscpFfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAscpFfoQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.ffo.query"
}

func (r AliexpressAscpFfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAscpFfoQueryRequest) SetFulfillmentForwardOrderQuery(fulfillmentForwardOrderQuery *FulfillmentForwardOrderQueryDto) error {
    r.fulfillmentForwardOrderQuery = fulfillmentForwardOrderQuery
    r.Set("fulfillment_forward_order_query", fulfillmentForwardOrderQuery)
    return nil
}

func (r AliexpressAscpFfoQueryRequest) GetFulfillmentForwardOrderQuery() *FulfillmentForwardOrderQueryDto {
    return r.fulfillmentForwardOrderQuery
}

