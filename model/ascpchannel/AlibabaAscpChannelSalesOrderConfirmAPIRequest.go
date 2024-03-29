package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSalesOrderConfirmAPIRequest 渠道销售单确认收货 API请求
// alibaba.ascp.channel.sales.order.confirm
//
// 渠道销售单确认收货
type AlibabaAscpChannelSalesOrderConfirmAPIRequest struct {
	model.Params
	// 请求参数
	_confirmOrderRequest *ExternalConfirmSalesOrderRequest
}

// NewAlibabaAscpChannelSalesOrderConfirmRequest 初始化AlibabaAscpChannelSalesOrderConfirmAPIRequest对象
func NewAlibabaAscpChannelSalesOrderConfirmRequest() *AlibabaAscpChannelSalesOrderConfirmAPIRequest {
	return &AlibabaAscpChannelSalesOrderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelSalesOrderConfirmAPIRequest) Reset() {
	r._confirmOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSalesOrderConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.sales.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSalesOrderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelSalesOrderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirmOrderRequest is ConfirmOrderRequest Setter
// 请求参数
func (r *AlibabaAscpChannelSalesOrderConfirmAPIRequest) SetConfirmOrderRequest(_confirmOrderRequest *ExternalConfirmSalesOrderRequest) error {
	r._confirmOrderRequest = _confirmOrderRequest
	r.Set("confirm_order_request", _confirmOrderRequest)
	return nil
}

// GetConfirmOrderRequest ConfirmOrderRequest Getter
func (r AlibabaAscpChannelSalesOrderConfirmAPIRequest) GetConfirmOrderRequest() *ExternalConfirmSalesOrderRequest {
	return r._confirmOrderRequest
}

var poolAlibabaAscpChannelSalesOrderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelSalesOrderConfirmRequest()
	},
}

// GetAlibabaAscpChannelSalesOrderConfirmRequest 从 sync.Pool 获取 AlibabaAscpChannelSalesOrderConfirmAPIRequest
func GetAlibabaAscpChannelSalesOrderConfirmAPIRequest() *AlibabaAscpChannelSalesOrderConfirmAPIRequest {
	return poolAlibabaAscpChannelSalesOrderConfirmAPIRequest.Get().(*AlibabaAscpChannelSalesOrderConfirmAPIRequest)
}

// ReleaseAlibabaAscpChannelSalesOrderConfirmAPIRequest 将 AlibabaAscpChannelSalesOrderConfirmAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelSalesOrderConfirmAPIRequest(v *AlibabaAscpChannelSalesOrderConfirmAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelSalesOrderConfirmAPIRequest.Put(v)
}
