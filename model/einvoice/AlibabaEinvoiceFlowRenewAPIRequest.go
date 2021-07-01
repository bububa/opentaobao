package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceFlowRenewAPIRequest
工单(入驻、加盘、续约)续约 API请求
alibaba.einvoice.flow.renew

工单(含入驻、加盘、续约工单)续约能力开放 */
type AlibabaEinvoiceFlowRenewAPIRequest struct {
	model.Params
	// 续约请求参数
	_invoiceFlowRenewDto *InvoiceFlowRenewDto
}

// New
