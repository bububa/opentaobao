package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelOrderStatusUpdateAPIRequest 订单状态变更 API请求
// alibaba.wdk.channel.order.status.update
//
// 订单状态变更
type AlibabaWdkChannelOrderStatusUpdateAPIRequest struct {
	model.Params
	// 修改信息
	_orderStatusInfo *OrderStatusInfo
}

// NewAlibabaWdkChannelOrderStatusUpdateRequest 初始化AlibabaWdkChannelOrderStatusUpdateAPIRequest对象
func NewAlibabaWdkChannelOrderStatusUpdateRequest() *AlibabaWdkChannelOrderStatusUpdateAPIRequest {
	return &AlibabaWdkChannelOrderStatusUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkChannelOrderStatusUpdateAPIRequest) Reset() {
	r._orderStatusInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkChannelOrderStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderStatusInfo is OrderStatusInfo Setter
// 修改信息
func (r *AlibabaWdkChannelOrderStatusUpdateAPIRequest) SetOrderStatusInfo(_orderStatusInfo *OrderStatusInfo) error {
	r._orderStatusInfo = _orderStatusInfo
	r.Set("order_status_info", _orderStatusInfo)
	return nil
}

// GetOrderStatusInfo OrderStatusInfo Getter
func (r AlibabaWdkChannelOrderStatusUpdateAPIRequest) GetOrderStatusInfo() *OrderStatusInfo {
	return r._orderStatusInfo
}

var poolAlibabaWdkChannelOrderStatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkChannelOrderStatusUpdateRequest()
	},
}

// GetAlibabaWdkChannelOrderStatusUpdateRequest 从 sync.Pool 获取 AlibabaWdkChannelOrderStatusUpdateAPIRequest
func GetAlibabaWdkChannelOrderStatusUpdateAPIRequest() *AlibabaWdkChannelOrderStatusUpdateAPIRequest {
	return poolAlibabaWdkChannelOrderStatusUpdateAPIRequest.Get().(*AlibabaWdkChannelOrderStatusUpdateAPIRequest)
}

// ReleaseAlibabaWdkChannelOrderStatusUpdateAPIRequest 将 AlibabaWdkChannelOrderStatusUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkChannelOrderStatusUpdateAPIRequest(v *AlibabaWdkChannelOrderStatusUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkChannelOrderStatusUpdateAPIRequest.Put(v)
}
