package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsalesordercreateAPIRequest 供应链渠道销售单创建接口 API请求
// alibaba.ascp.channel.sales.order.create
//
// 阿里巴巴供应链渠道销售订单创建接口
type AlibabaascpchannelsalesordercreateAPIRequest struct {
	model.Params
	// 请求参数
	_createOrderRequest *ExternalCreateSalesOrderRequest
}

// NewAlibabaascpchannelsalesordercreateRequest 初始化AlibabaascpchannelsalesordercreateAPIRequest对象
func NewAlibabaascpchannelsalesordercreateRequest() *AlibabaascpchannelsalesordercreateAPIRequest {
	return &AlibabaascpchannelsalesordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelsalesordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.sales.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelsalesordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelsalesordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateOrderRequest is CreateOrderRequest Setter
// 请求参数
func (r *AlibabaascpchannelsalesordercreateAPIRequest) SetCreateOrderRequest(_createOrderRequest *ExternalCreateSalesOrderRequest) error {
	r._createOrderRequest = _createOrderRequest
	r.Set("create_order_request", _createOrderRequest)
	return nil
}

// GetCreateOrderRequest CreateOrderRequest Getter
func (r AlibabaascpchannelsalesordercreateAPIRequest) GetCreateOrderRequest() *ExternalCreateSalesOrderRequest {
	return r._createOrderRequest
}
