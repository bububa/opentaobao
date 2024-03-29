package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest 供应链外部订单创建前校验接口 API请求
// alibaba.ascp.channel.sales.order.before.check
//
// 供应链外部订单创建前校验接口
type AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest struct {
	model.Params
	// 请求参数
	_orderCheckRequest *ExtOrderCheckRequest
}

// NewAlibabaAscpChannelSalesOrderBeforeCheckRequest 初始化AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest对象
func NewAlibabaAscpChannelSalesOrderBeforeCheckRequest() *AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest {
	return &AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest) Reset() {
	r._orderCheckRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.sales.order.before.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCheckRequest is OrderCheckRequest Setter
// 请求参数
func (r *AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest) SetOrderCheckRequest(_orderCheckRequest *ExtOrderCheckRequest) error {
	r._orderCheckRequest = _orderCheckRequest
	r.Set("order_check_request", _orderCheckRequest)
	return nil
}

// GetOrderCheckRequest OrderCheckRequest Getter
func (r AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest) GetOrderCheckRequest() *ExtOrderCheckRequest {
	return r._orderCheckRequest
}

var poolAlibabaAscpChannelSalesOrderBeforeCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelSalesOrderBeforeCheckRequest()
	},
}

// GetAlibabaAscpChannelSalesOrderBeforeCheckRequest 从 sync.Pool 获取 AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest
func GetAlibabaAscpChannelSalesOrderBeforeCheckAPIRequest() *AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest {
	return poolAlibabaAscpChannelSalesOrderBeforeCheckAPIRequest.Get().(*AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest)
}

// ReleaseAlibabaAscpChannelSalesOrderBeforeCheckAPIRequest 将 AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelSalesOrderBeforeCheckAPIRequest(v *AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelSalesOrderBeforeCheckAPIRequest.Put(v)
}
