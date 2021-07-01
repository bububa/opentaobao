package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceOrderRefundUpdateAPIRequest
回传订单退款审核结果 API请求
alibaba.einvoice.order.refund.update

ISV回传订单退款审核结果 */
type AlibabaEinvoiceOrderRefundUpdateAPIRequest struct {
	model.Params
	// 退款审核结果DTO
	_orderRefundResultDto *InvoiceOrderRefundResultDto
}

// New
