package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyRefundDisagreeAPIRequest
saas 售后逆向 商户拒绝用户逆向申请 API请求
alibaba.tcls.aelophy.refund.disagree

saas 售后逆向 商户拒绝用户逆向申请 */
type AlibabaTclsAelophyRefundDisagreeAPIRequest struct {
	model.Params
	// 退款单ID
	_refundId string
	// 拒绝原因
	_rejectReason string
}

// New
