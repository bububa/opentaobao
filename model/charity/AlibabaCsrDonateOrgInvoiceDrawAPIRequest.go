package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrDonateOrgInvoiceDrawAPIRequest 机构开具商家票据信息 API请求
// alibaba.csr.donate.org.invoice.draw
//
// 机构开具商家票据信息
type AlibabaCsrDonateOrgInvoiceDrawAPIRequest struct {
	model.Params
	// 参数
	_csrInvoiceExternalOrgDrawDto *CsrInvoiceExternalOrgDrawDto
}

// NewAlibabaCsrDonateOrgInvoiceDrawRequest 初始化AlibabaCsrDonateOrgInvoiceDrawAPIRequest对象
func NewAlibabaCsrDonateOrgInvoiceDrawRequest() *AlibabaCsrDonateOrgInvoiceDrawAPIRequest {
	return &AlibabaCsrDonateOrgInvoiceDrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCsrDonateOrgInvoiceDrawAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.donate.org.invoice.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCsrDonateOrgInvoiceDrawAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCsrDonateOrgInvoiceDrawAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCsrInvoiceExternalOrgDrawDto is CsrInvoiceExternalOrgDrawDto Setter
// 参数
func (r *AlibabaCsrDonateOrgInvoiceDrawAPIRequest) SetCsrInvoiceExternalOrgDrawDto(_csrInvoiceExternalOrgDrawDto *CsrInvoiceExternalOrgDrawDto) error {
	r._csrInvoiceExternalOrgDrawDto = _csrInvoiceExternalOrgDrawDto
	r.Set("csr_invoice_external_org_draw_dto", _csrInvoiceExternalOrgDrawDto)
	return nil
}

// GetCsrInvoiceExternalOrgDrawDto CsrInvoiceExternalOrgDrawDto Getter
func (r AlibabaCsrDonateOrgInvoiceDrawAPIRequest) GetCsrInvoiceExternalOrgDrawDto() *CsrInvoiceExternalOrgDrawDto {
	return r._csrInvoiceExternalOrgDrawDto
}
