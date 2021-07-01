package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest
翱象商家自有渠道 订单状态更新 API请求
alibaba.tcls.aelophy.merchant.channel.order.updatestatus

订单状态变更 */
type AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest struct {
	model.Params
	// 修改信息
	_orderStatusInfo *OrderStatusInfo
}

// New
