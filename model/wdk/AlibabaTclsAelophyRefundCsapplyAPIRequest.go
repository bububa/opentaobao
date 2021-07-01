package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyRefundCsapplyAPIRequest
商家代客售后提交逆向申请 API请求
alibaba.tcls.aelophy.refund.csapply

商家代客售后提交逆向申请 */
type AlibabaTclsAelophyRefundCsapplyAPIRequest struct {
	model.Params
	// 逆向申请入参
	_refundCsApplyDTO *RefundCsApplyDto
}

// New
