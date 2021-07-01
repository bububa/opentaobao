package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest
翱象商家自有渠道 逆向单申请 API请求
alibaba.tcls.aelophy.merchant.channel.refund.apply

翱象小程序 用户逆向单申请 */
type AlibabaTclsAelophyMerchantChannelRefundApplyAPIRequest struct {
	model.Params
	// 请求对象
	_refundApplyInfo *RefundApplyInfo
}

// New
