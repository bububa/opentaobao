package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest
翱象商家自有渠道 交易订单取消 API请求
alibaba.tcls.aelophy.merchant.channel.order.cancel

翱象小程序用户取消订单 */
type AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest struct {
	model.Params
	// 取消信息
	_userCancelInfo *OrderUserCancelInfo
}

// New
