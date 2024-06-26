package ascpchannel

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelSalesOrderCreateAPIRequest) Reset() {
	r._createOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSalesOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.sales.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSalesOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelSalesOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateOrderRequest is CreateOrderRequest Setter
// 请求参数
func (r *AlibabaAscpChannelSalesOrderCreateAPIRequest) SetCreateOrderRequest(_createOrderRequest *ExternalCreateSalesOrderRequest) error {
	r._createOrderRequest = _createOrderRequest
	r.Set("create_order_request", _createOrderRequest)
	return nil
}

// GetCreateOrderRequest CreateOrderRequest Getter
func (r AlibabaAscpChannelSalesOrderCreateAPIRequest) GetCreateOrderRequest() *ExternalCreateSalesOrderRequest {
	return r._createOrderRequest
}

var poolAlibabaAscpChannelSalesOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelSalesOrderCreateRequest()
	},
}

// GetAlibabaAscpChannelSalesOrderCreateRequest 从 sync.Pool 获取 AlibabaAscpChannelSalesOrderCreateAPIRequest
func GetAlibabaAscpChannelSalesOrderCreateAPIRequest() *AlibabaAscpChannelSalesOrderCreateAPIRequest {
	return poolAlibabaAscpChannelSalesOrderCreateAPIRequest.Get().(*AlibabaAscpChannelSalesOrderCreateAPIRequest)
}

// ReleaseAlibabaAscpChannelSalesOrderCreateAPIRequest 将 AlibabaAscpChannelSalesOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelSalesOrderCreateAPIRequest(v *AlibabaAscpChannelSalesOrderCreateAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelSalesOrderCreateAPIRequest.Put(v)
}
