package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkchannelorderstatusupdateAPIRequest 订单状态变更 API请求
// alibaba.wdk.channel.order.status.update
//
// 订单状态变更
type AlibabawdkchannelorderstatusupdateAPIRequest struct {
	model.Params
	// 修改信息
	_orderStatusInfo *OrderStatusInfo
}

// NewAlibabawdkchannelorderstatusupdateRequest 初始化AlibabawdkchannelorderstatusupdateAPIRequest对象
func NewAlibabawdkchannelorderstatusupdateRequest() *AlibabawdkchannelorderstatusupdateAPIRequest {
	return &AlibabawdkchannelorderstatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkchannelorderstatusupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkchannelorderstatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkchannelorderstatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderStatusInfo is OrderStatusInfo Setter
// 修改信息
func (r *AlibabawdkchannelorderstatusupdateAPIRequest) SetOrderStatusInfo(_orderStatusInfo *OrderStatusInfo) error {
	r._orderStatusInfo = _orderStatusInfo
	r.Set("order_status_info", _orderStatusInfo)
	return nil
}

// GetOrderStatusInfo OrderStatusInfo Getter
func (r AlibabawdkchannelorderstatusupdateAPIRequest) GetOrderStatusInfo() *OrderStatusInfo {
	return r._orderStatusInfo
}
