package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress销退单明细查询API APIRequest
aliexpress.ascp.fro.item.query

AE履约销退单明细查询API
*/
type AliexpressAscpFroItemQueryRequest struct {
    model.Params

    // dto
    fulfillmentReverseOrderItemQuery   *FulfillmentReverseOrderItemQueryDto 

}

func NewAliexpressAscpFroItemQueryRequest() *AliexpressAscpFroItemQueryRequest{
    return &AliexpressAscpFroItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAscpFroItemQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.fro.item.query"
}

func (r AliexpressAscpFroItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAscpFroItemQueryRequest) SetFulfillmentReverseOrderItemQuery(fulfillmentReverseOrderItemQuery *FulfillmentReverseOrderItemQueryDto) error {
    r.fulfillmentReverseOrderItemQuery = fulfillmentReverseOrderItemQuery
    r.Set("fulfillment_reverse_order_item_query", fulfillmentReverseOrderItemQuery)
    return nil
}

func (r AliexpressAscpFroItemQueryRequest) GetFulfillmentReverseOrderItemQuery() *FulfillmentReverseOrderItemQueryDto {
    return r.fulfillmentReverseOrderItemQuery
}

