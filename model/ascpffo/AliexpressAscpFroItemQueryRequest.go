package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress销退单明细查询API API请求
aliexpress.ascp.fro.item.query

AE履约销退单明细查询API
*/
type AliexpressAscpFroItemQueryRequest struct {
    model.Params
    // dto
    fulfillmentReverseOrderItemQuery   *FulfillmentReverseOrderItemQueryDto
}

// 初始化AliexpressAscpFroItemQueryRequest对象
func NewAliexpressAscpFroItemQueryRequest() *AliexpressAscpFroItemQueryRequest{
    return &AliexpressAscpFroItemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpFroItemQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.fro.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpFroItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FulfillmentReverseOrderItemQuery Setter
// dto
func (r *AliexpressAscpFroItemQueryRequest) SetFulfillmentReverseOrderItemQuery(fulfillmentReverseOrderItemQuery *FulfillmentReverseOrderItemQueryDto) error {
    r.fulfillmentReverseOrderItemQuery = fulfillmentReverseOrderItemQuery
    r.Set("fulfillment_reverse_order_item_query", fulfillmentReverseOrderItemQuery)
    return nil
}

// FulfillmentReverseOrderItemQuery Getter
func (r AliexpressAscpFroItemQueryRequest) GetFulfillmentReverseOrderItemQuery() *FulfillmentReverseOrderItemQueryDto {
    return r.fulfillmentReverseOrderItemQuery
}
