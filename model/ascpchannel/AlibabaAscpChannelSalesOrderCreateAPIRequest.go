package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSalesOrderCreateAPIRequest 供应链渠道销售单创建接口 API请求
// alibaba.ascp.channel.sales.order.create
//
// 阿里巴巴供应链渠道销售订单创建接口
type AlibabaAscpChannelSalesOrderCreateAPIRequest struct {
	model.Params
	// 请求参数
	_createOrderRequest *ExternalCreateSalesOrderRequest
}

// NewAlibabaAscpChannelSalesOrderCreateRequest 初始化AlibabaAscpChannelSalesOrderCreateAPIRequest对象
func NewAlibabaAscpChannelSalesOrderCreateRequest() *AlibabaAscpChannelSalesOrderCreateAPIRequest {
	return &AlibabaAscpChannelSalesOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSalesOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.sales.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSalesOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CreateOrderRequest Setter
// 请求参数
func (r *AlibabaAscpChannelSalesOrderCreateAPIRequest) SetCreateOrderRequest(_createOrderRequest *ExternalCreateSalesOrderRequest) error {
	r._createOrderRequest = _createOrderRequest
	r.Set("create_order_request", _createOrderRequest)
	return nil
}

// Get CreateOrderRequest Getter
func (r AlibabaAscpChannelSalesOrderCreateAPIRequest) GetCreateOrderRequest() *ExternalCreateSalesOrderRequest {
	return r._createOrderRequest
}
