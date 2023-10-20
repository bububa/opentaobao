package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascpffoqueryAPIRequest AliExpress发货单查询API API请求
// aliexpress.ascp.ffo.query
//
// AE 履约发货单分页查询接口
type AliexpressascpffoqueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentForwardOrderQuery *FulfillmentForwardOrderQueryDto
}

// NewAliexpressascpffoqueryRequest 初始化AliexpressascpffoqueryAPIRequest对象
func NewAliexpressascpffoqueryRequest() *AliexpressascpffoqueryAPIRequest {
	return &AliexpressascpffoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressascpffoqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.ascp.ffo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressascpffoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressascpffoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillmentForwardOrderQuery is FulfillmentForwardOrderQuery Setter
// dto
func (r *AliexpressascpffoqueryAPIRequest) SetFulfillmentForwardOrderQuery(_fulfillmentForwardOrderQuery *FulfillmentForwardOrderQueryDto) error {
	r._fulfillmentForwardOrderQuery = _fulfillmentForwardOrderQuery
	r.Set("fulfillment_forward_order_query", _fulfillmentForwardOrderQuery)
	return nil
}

// GetFulfillmentForwardOrderQuery FulfillmentForwardOrderQuery Getter
func (r AliexpressascpffoqueryAPIRequest) GetFulfillmentForwardOrderQuery() *FulfillmentForwardOrderQueryDto {
	return r._fulfillmentForwardOrderQuery
}
