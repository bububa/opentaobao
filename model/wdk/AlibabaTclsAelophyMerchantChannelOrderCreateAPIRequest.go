package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest
翱象商家自有渠道 订单创建 API请求
alibaba.tcls.aelophy.merchant.channel.order.create

翱象小程序渠道订单创建 */
type AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest struct {
	model.Params
	// 订单信息
	_orderInfo *OrderInfo
}

// New
