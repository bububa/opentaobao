package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkChannelOrderStatusUpdateAPIRequest
订单状态变更 API请求
alibaba.wdk.channel.order.status.update

订单状态变更 */
type AlibabaWdkChannelOrderStatusUpdateAPIRequest struct {
	model.Params
	// 修改信息
	_orderStatusInfo *OrderStatusInfo
}

// New
