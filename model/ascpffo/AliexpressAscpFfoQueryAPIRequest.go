package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAscpFfoQueryAPIRequest AliExpress发货单查询API API请求
// aliexpress.ascp.ffo.query
//
// AE 履约发货单分页查询接口
type AliexpressAscpFfoQueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentForwardOrderQuery *FulfillmentForwardOrderQueryDto
}

// NewAliexpressAscpFfoQueryRequest 初始化AliexpressAscpFfoQueryAPIRequest对象
func NewAliexpressAscpFfoQueryRequest() *AliexpressAscpFfoQueryAPIRequest {
	return &AliexpressAscpFfoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAscpFfoQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.ffo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAscpFfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAscpFfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillmentForwardOrderQuery is FulfillmentForwardOrderQuery Setter
// dto
func (r *AliexpressAscpFfoQueryAPIRequest) SetFulfillmentForwardOrderQuery(_fulfillmentForwardOrderQuery *FulfillmentForwardOrderQueryDto) error {
	r._fulfillmentForwardOrderQuery = _fulfillmentForwardOrderQuery
	r.Set("fulfillment_forward_order_query", _fulfillmentForwardOrderQuery)
	return nil
}

// GetFulfillmentForwardOrderQuery FulfillmentForwardOrderQuery Getter
func (r AliexpressAscpFfoQueryAPIRequest) GetFulfillmentForwardOrderQuery() *FulfillmentForwardOrderQueryDto {
	return r._fulfillmentForwardOrderQuery
}
