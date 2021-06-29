package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress发货单明细分页查询API API请求
aliexpress.ascp.ffo.item.query

AE履约发货单明细分页查询
*/
type AliexpressAscpFfoItemQueryRequest struct {
    model.Params
    // DTO
    _fulfillmentForwardOrderItemQuery   *FulfillmentForwardOrderItemQueryDto
}

// 初始化AliexpressAscpFfoItemQueryRequest对象
func NewAliexpressAscpFfoItemQueryRequest() *AliexpressAscpFfoItemQueryRequest{
    return &AliexpressAscpFfoItemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpFfoItemQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.ffo.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpFfoItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FulfillmentForwardOrderItemQuery Setter
// DTO
func (r *AliexpressAscpFfoItemQueryRequest) SetFulfillmentForwardOrderItemQuery(_fulfillmentForwardOrderItemQuery *FulfillmentForwardOrderItemQueryDto) error {
    r._fulfillmentForwardOrderItemQuery = _fulfillmentForwardOrderItemQuery
    r.Set("fulfillment_forward_order_item_query", _fulfillmentForwardOrderItemQuery)
    return nil
}

// FulfillmentForwardOrderItemQuery Getter
func (r AliexpressAscpFfoItemQueryRequest) GetFulfillmentForwardOrderItemQuery() *FulfillmentForwardOrderItemQueryDto {
    return r._fulfillmentForwardOrderItemQuery
}
