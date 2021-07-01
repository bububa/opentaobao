package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceInvoiceapplyGetAPIRequest
获取商家的开票申请 API请求
alibaba.einvoice.invoiceapply.get

开票服务商接收到商家发起的开票申请消息后，调用此接口拉取商家详细的开票申请内容 */
type AlibabaEinvoiceInvoiceapplyGetAPIRequest struct {
	model.Params
	// 开票申请id
	_applyId string
}

// New
