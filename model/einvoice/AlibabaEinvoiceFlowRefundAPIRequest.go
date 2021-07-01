package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceFlowRefundAPIRequest
退订工单(入驻、加盘、续约) API请求
alibaba.einvoice.flow.refund

电子发票工单系统，工单退订能力开放 */
type AlibabaEinvoiceFlowRefundAPIRequest struct {
	model.Params
	// 退订请求参数
	_invoiceFlowRefund *InvoiceFlowRefundDto
}

// New
