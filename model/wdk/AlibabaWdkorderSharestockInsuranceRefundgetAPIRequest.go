package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest
共享库存投保业务售后逆向订单数据获取 API请求
alibaba.wdkorder.sharestock.insurance.refundget

共享库存投保业务售后逆向订单数据获取 */
type AlibabaWdkorderSharestockInsuranceRefundgetAPIRequest struct {
	model.Params
	// 淘宝子订单ID
	_tbSubOrderId string
	// 退货单ID
	_refundId string
}

// New
