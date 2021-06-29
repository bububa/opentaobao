package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress销退单查询API API请求
aliexpress.ascp.fro.query

AE履约销退单查询接口
*/
type AliexpressAscpFroQueryRequest struct {
    model.Params
    // dto
    fulfillmentReverseOrderQuery   *FulfillmentReverseOrderQueryDto
}

// 初始化AliexpressAscpFroQueryRequest对象
func NewAliexpressAscpFroQueryRequest() *AliexpressAscpFroQueryRequest{
    return &AliexpressAscpFroQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpFroQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.fro.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpFroQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FulfillmentReverseOrderQuery Setter
// dto
func (r *AliexpressAscpFroQueryRequest) SetFulfillmentReverseOrderQuery(fulfillmentReverseOrderQuery *FulfillmentReverseOrderQueryDto) error {
    r.fulfillmentReverseOrderQuery = fulfillmentReverseOrderQuery
    r.Set("fulfillment_reverse_order_query", fulfillmentReverseOrderQuery)
    return nil
}

// FulfillmentReverseOrderQuery Getter
func (r AliexpressAscpFroQueryRequest) GetFulfillmentReverseOrderQuery() *FulfillmentReverseOrderQueryDto {
    return r.fulfillmentReverseOrderQuery
}
