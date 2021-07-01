package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceProdApplyAPIRequest
提交发票申请 API请求
alibaba.einvoice.prod.apply

提交开票申请，如果商户授权自动开票则自动转开票，否则等待商户审核。 */
type AlibabaEinvoiceProdApplyAPIRequest struct {
	model.Params
	// 申请开票请求
	_paramInvoiceApplyDto *InvoiceApplyDto
}

// New
