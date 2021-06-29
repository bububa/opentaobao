package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress发货单查询API API请求
aliexpress.ascp.ffo.query

AE 履约发货单分页查询接口
*/
type AliexpressAscpFfoQueryRequest struct {
    model.Params
    // dto
    fulfillmentForwardOrderQuery   *FulfillmentForwardOrderQueryDto
}

// 初始化AliexpressAscpFfoQueryRequest对象
func NewAliexpressAscpFfoQueryRequest() *AliexpressAscpFfoQueryRequest{
    return &AliexpressAscpFfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpFfoQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.ffo.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpFfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FulfillmentForwardOrderQuery Setter
// dto
func (r *AliexpressAscpFfoQueryRequest) SetFulfillmentForwardOrderQuery(fulfillmentForwardOrderQuery *FulfillmentForwardOrderQueryDto) error {
    r.fulfillmentForwardOrderQuery = fulfillmentForwardOrderQuery
    r.Set("fulfillment_forward_order_query", fulfillmentForwardOrderQuery)
    return nil
}

// FulfillmentForwardOrderQuery Getter
func (r AliexpressAscpFfoQueryRequest) GetFulfillmentForwardOrderQuery() *FulfillmentForwardOrderQueryDto {
    return r.fulfillmentForwardOrderQuery
}
