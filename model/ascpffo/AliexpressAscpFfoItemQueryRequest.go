package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress发货单明细分页查询API APIRequest
aliexpress.ascp.ffo.item.query

AE履约发货单明细分页查询
*/
type AliexpressAscpFfoItemQueryRequest struct {
    model.Params

    // DTO
    fulfillmentForwardOrderItemQuery   *FulfillmentForwardOrderItemQueryDto 

}

func NewAliexpressAscpFfoItemQueryRequest() *AliexpressAscpFfoItemQueryRequest{
    return &AliexpressAscpFfoItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAscpFfoItemQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.ffo.item.query"
}

func (r AliexpressAscpFfoItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAscpFfoItemQueryRequest) SetFulfillmentForwardOrderItemQuery(fulfillmentForwardOrderItemQuery *FulfillmentForwardOrderItemQueryDto) error {
    r.fulfillmentForwardOrderItemQuery = fulfillmentForwardOrderItemQuery
    r.Set("fulfillment_forward_order_item_query", fulfillmentForwardOrderItemQuery)
    return nil
}

func (r AliexpressAscpFfoItemQueryRequest) GetFulfillmentForwardOrderItemQuery() *FulfillmentForwardOrderItemQueryDto {
    return r.fulfillmentForwardOrderItemQuery
}

