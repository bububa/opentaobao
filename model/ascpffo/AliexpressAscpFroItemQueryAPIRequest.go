package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascpfroitemqueryAPIRequest AliExpress销退单明细查询API API请求
// aliexpress.ascp.fro.item.query
//
// AE履约销退单明细查询API
type AliexpressascpfroitemqueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentReverseOrderItemQuery *FulfillmentReverseOrderItemQueryDto
}

// NewAliexpressascpfroitemqueryRequest 初始化AliexpressascpfroitemqueryAPIRequest对象
func NewAliexpressascpfroitemqueryRequest() *AliexpressascpfroitemqueryAPIRequest {
	return &AliexpressascpfroitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascpfroitemqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.fro.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascpfroitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascpfroitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillmentReverseOrderItemQuery is FulfillmentReverseOrderItemQuery Setter
// dto
func (r *AliexpressascpfroitemqueryAPIRequest) SetFulfillmentReverseOrderItemQuery(_fulfillmentReverseOrderItemQuery *FulfillmentReverseOrderItemQueryDto) error {
	r._fulfillmentReverseOrderItemQuery = _fulfillmentReverseOrderItemQuery
	r.Set("fulfillment_reverse_order_item_query", _fulfillmentReverseOrderItemQuery)
	return nil
}

// GetFulfillmentReverseOrderItemQuery FulfillmentReverseOrderItemQuery Getter
func (r AliexpressascpfroitemqueryAPIRequest) GetFulfillmentReverseOrderItemQuery() *FulfillmentReverseOrderItemQueryDto {
	return r._fulfillmentReverseOrderItemQuery
}
