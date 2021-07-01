package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpFroItemQueryAPIRequest
AliExpress销退单明细查询API API请求
aliexpress.ascp.fro.item.query

AE履约销退单明细查询API */
type AliexpressAscpFroItemQueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentReverseOrderItemQuery *FulfillmentReverseOrderItemQueryDto
}

// NewAliexpressAscpFroItemQueryRequest 初始化AliexpressAscpFroItemQueryAPIRequest对象
func NewAliexpressAscpFroItemQueryRequest() *AliexpressAscpFroItemQueryAPIRequest {
	return &AliexpressAscpFroItemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpFroItemQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.fro.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpFroItemQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FulfillmentReverseOrderItemQuery Setter
// dto
func (r *AliexpressAscpFroItemQueryAPIRequest) SetFulfillmentReverseOrderItemQuery(_fulfillmentReverseOrderItemQuery *FulfillmentReverseOrderItemQueryDto) error {
	r._fulfillmentReverseOrderItemQuery = _fulfillmentReverseOrderItemQuery
	r.Set("fulfillment_reverse_order_item_query", _fulfillmentReverseOrderItemQuery)
	return nil
}

// Get FulfillmentReverseOrderItemQuery Getter
func (r AliexpressAscpFroItemQueryAPIRequest) GetFulfillmentReverseOrderItemQuery() *FulfillmentReverseOrderItemQueryDto {
	return r._fulfillmentReverseOrderItemQuery
}
