package charity

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCsrDonateOrgInvoiceDrawAPIRequest) Reset() {
	r._csrInvoiceExternalOrgDrawDto = nil
	r.Params.ToZero()
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

var poolAlibabaCsrDonateOrgInvoiceDrawAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCsrDonateOrgInvoiceDrawRequest()
	},
}

// GetAlibabaCsrDonateOrgInvoiceDrawRequest 从 sync.Pool 获取 AlibabaCsrDonateOrgInvoiceDrawAPIRequest
func GetAlibabaCsrDonateOrgInvoiceDrawAPIRequest() *AlibabaCsrDonateOrgInvoiceDrawAPIRequest {
	return poolAlibabaCsrDonateOrgInvoiceDrawAPIRequest.Get().(*AlibabaCsrDonateOrgInvoiceDrawAPIRequest)
}

// ReleaseAlibabaCsrDonateOrgInvoiceDrawAPIRequest 将 AlibabaCsrDonateOrgInvoiceDrawAPIRequest 放入 sync.Pool
func ReleaseAlibabaCsrDonateOrgInvoiceDrawAPIRequest(v *AlibabaCsrDonateOrgInvoiceDrawAPIRequest) {
	v.Reset()
	poolAlibabaCsrDonateOrgInvoiceDrawAPIRequest.Put(v)
}
