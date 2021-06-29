package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress销退单查询API APIRequest
aliexpress.ascp.fro.query

AE履约销退单查询接口
*/
type AliexpressAscpFroQueryRequest struct {
    model.Params

    // dto
    fulfillmentReverseOrderQuery   *FulfillmentReverseOrderQueryDto 

}

func NewAliexpressAscpFroQueryRequest() *AliexpressAscpFroQueryRequest{
    return &AliexpressAscpFroQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAscpFroQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.fro.query"
}

func (r AliexpressAscpFroQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAscpFroQueryRequest) SetFulfillmentReverseOrderQuery(fulfillmentReverseOrderQuery *FulfillmentReverseOrderQueryDto) error {
    r.fulfillmentReverseOrderQuery = fulfillmentReverseOrderQuery
    r.Set("fulfillment_reverse_order_query", fulfillmentReverseOrderQuery)
    return nil
}

func (r AliexpressAscpFroQueryRequest) GetFulfillmentReverseOrderQuery() *FulfillmentReverseOrderQueryDto {
    return r.fulfillmentReverseOrderQuery
}

