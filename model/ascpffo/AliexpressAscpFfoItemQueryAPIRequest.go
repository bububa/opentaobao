package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpFfoItemQueryAPIRequest AliExpress发货单明细分页查询API API请求
// aliexpress.ascp.ffo.item.query
//
// AE履约发货单明细分页查询
type AliexpressAscpFfoItemQueryAPIRequest struct {
	model.Params
	// DTO
	_fulfillmentForwardOrderItemQuery *FulfillmentForwardOrderItemQueryDto
}

// NewAliexpressAscpFfoItemQueryRequest 初始化AliexpressAscpFfoItemQueryAPIRequest对象
func NewAliexpressAscpFfoItemQueryRequest() *AliexpressAscpFfoItemQueryAPIRequest {
	return &AliexpressAscpFfoItemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpFfoItemQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.ffo.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpFfoItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpFfoItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillmentForwardOrderItemQuery is FulfillmentForwardOrderItemQuery Setter
// DTO
func (r *AliexpressAscpFfoItemQueryAPIRequest) SetFulfillmentForwardOrderItemQuery(_fulfillmentForwardOrderItemQuery *FulfillmentForwardOrderItemQueryDto) error {
	r._fulfillmentForwardOrderItemQuery = _fulfillmentForwardOrderItemQuery
	r.Set("fulfillment_forward_order_item_query", _fulfillmentForwardOrderItemQuery)
	return nil
}

// GetFulfillmentForwardOrderItemQuery FulfillmentForwardOrderItemQuery Getter
func (r AliexpressAscpFfoItemQueryAPIRequest) GetFulfillmentForwardOrderItemQuery() *FulfillmentForwardOrderItemQueryDto {
	return r._fulfillmentForwardOrderItemQuery
}
