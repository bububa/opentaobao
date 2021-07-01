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
type AliexpressAscpFfoItemQueryAPIRequest struct {
    model.Params
    // DTO
    _fulfillmentForwardOrderItemQuery   *FulfillmentForwardOrderItemQueryDTO
}

// 初始化AliexpressAscpFfoItemQueryAPIRequest对象
func NewAliexpressAscpFfoItemQueryRequest() *AliexpressAscpFfoItemQueryAPIRequest{
    return &AliexpressAscpFfoItemQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpFfoItemQueryAPIRequest) GetApiMethodName() string {
    return "aliexpress.ascp.ffo.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpFfoItemQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FulfillmentForwardOrderItemQuery Setter
// DTO
func (r *AliexpressAscpFfoItemQueryAPIRequest) SetFulfillmentForwardOrderItemQuery(_fulfillmentForwardOrderItemQuery *FulfillmentForwardOrderItemQueryDTO) error {
    r._fulfillmentForwardOrderItemQuery = _fulfillmentForwardOrderItemQuery
    r.Set("fulfillment_forward_order_item_query", _fulfillmentForwardOrderItemQuery)
    return nil
}

// FulfillmentForwardOrderItemQuery Getter
func (r AliexpressAscpFfoItemQueryAPIRequest) GetFulfillmentForwardOrderItemQuery() *FulfillmentForwardOrderItemQueryDTO {
    return r._fulfillmentForwardOrderItemQuery
}
