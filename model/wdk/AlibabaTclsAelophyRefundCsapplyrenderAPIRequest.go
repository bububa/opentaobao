package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyRefundCsapplyrenderAPIRequest
商家代客售后逆向申请渲染获取 API请求
alibaba.tcls.aelophy.refund.csapplyrender

提供商家代客售后逆向申请渲染获取的接口 */
type AlibabaTclsAelophyRefundCsapplyrenderAPIRequest struct {
	model.Params
	// 系统自动生成
	_refundCsApplyRenderDTO *RefundCsApplyRenderDto
}

// New
