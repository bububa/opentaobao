package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacsrdonateorginvoicerejectAPIRequest 机构驳回商家票据信息 API请求
// alibaba.csr.donate.org.invoice.reject
//
// 机构驳回商家票据信息
type AlibabacsrdonateorginvoicerejectAPIRequest struct {
	model.Params
	// 参数
	_csrInvoiceExternalOrgRejectDto *CsrInvoiceExternalOrgRejectDto
}

// NewAlibabacsrdonateorginvoicerejectRequest 初始化AlibabacsrdonateorginvoicerejectAPIRequest对象
func NewAlibabacsrdonateorginvoicerejectRequest() *AlibabacsrdonateorginvoicerejectAPIRequest {
	return &AlibabacsrdonateorginvoicerejectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacsrdonateorginvoicerejectAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.donate.org.invoice.reject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacsrdonateorginvoicerejectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacsrdonateorginvoicerejectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCsrInvoiceExternalOrgRejectDto is CsrInvoiceExternalOrgRejectDto Setter
// 参数
func (r *AlibabacsrdonateorginvoicerejectAPIRequest) SetCsrInvoiceExternalOrgRejectDto(_csrInvoiceExternalOrgRejectDto *CsrInvoiceExternalOrgRejectDto) error {
	r._csrInvoiceExternalOrgRejectDto = _csrInvoiceExternalOrgRejectDto
	r.Set("csr_invoice_external_org_reject_dto", _csrInvoiceExternalOrgRejectDto)
	return nil
}

// GetCsrInvoiceExternalOrgRejectDto CsrInvoiceExternalOrgRejectDto Getter
func (r AlibabacsrdonateorginvoicerejectAPIRequest) GetCsrInvoiceExternalOrgRejectDto() *CsrInvoiceExternalOrgRejectDto {
	return r._csrInvoiceExternalOrgRejectDto
}
