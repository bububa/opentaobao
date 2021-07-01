package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest
共享库存逆向订单理赔单回传 API请求
alibaba.wdkorder.sharestock.insurance.refundcallback

共享库存逆向订单理赔单回传 */
type AlibabaWdkorderSharestockInsuranceRefundcallbackAPIRequest struct {
	model.Params
	// 淘宝交易子单ID
	_tbSubOrderId int64
	// 退款单ID
	_refundId string
	// 理赔单ID
	_claimId string
}

// New
