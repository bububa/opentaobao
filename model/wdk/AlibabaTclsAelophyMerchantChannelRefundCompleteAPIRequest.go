package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest
翱象商家自有渠道 逆向单完成 API请求
alibaba.tcls.aelophy.merchant.channel.refund.complete

翱象小程序 退款完成 */
type AlibabaTclsAelophyMerchantChannelRefundCompleteAPIRequest struct {
	model.Params
	// 请求对象
	_refundCompleteInfo *RefundCompleteInfo
}

// New
