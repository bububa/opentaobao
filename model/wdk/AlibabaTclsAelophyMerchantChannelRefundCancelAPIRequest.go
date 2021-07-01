package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest
翱象商家自有渠道 逆向单申请取消 API请求
alibaba.tcls.aelophy.merchant.channel.refund.cancel

翱象小程序 用户逆向申请取消 */
type AlibabaTclsAelophyMerchantChannelRefundCancelAPIRequest struct {
	model.Params
	// 逆向申请取消
	_refundCancelInfo *RefundCancelInfo
}

// New
