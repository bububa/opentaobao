package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascpfroqueryAPIRequest AliExpress销退单查询API API请求
// aliexpress.ascp.fro.query
//
// AE履约销退单查询接口
type AliexpressascpfroqueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentReverseOrderQuery *FulfillmentReverseOrderQueryDto
}

// NewAliexpressascpfroqueryRequest 初始化AliexpressascpfroqueryAPIRequest对象
func NewAliexpressascpfroqueryRequest() *AliexpressascpfroqueryAPIRequest {
	return &AliexpressascpfroqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascpfroqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.fro.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascpfroqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascpfroqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillmentReverseOrderQuery is FulfillmentReverseOrderQuery Setter
// dto
func (r *AliexpressascpfroqueryAPIRequest) SetFulfillmentReverseOrderQuery(_fulfillmentReverseOrderQuery *FulfillmentReverseOrderQueryDto) error {
	r._fulfillmentReverseOrderQuery = _fulfillmentReverseOrderQuery
	r.Set("fulfillment_reverse_order_query", _fulfillmentReverseOrderQuery)
	return nil
}

// GetFulfillmentReverseOrderQuery FulfillmentReverseOrderQuery Getter
func (r AliexpressascpfroqueryAPIRequest) GetFulfillmentReverseOrderQuery() *FulfillmentReverseOrderQueryDto {
	return r._fulfillmentReverseOrderQuery
}
