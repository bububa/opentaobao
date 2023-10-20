package charity

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrDonateOrgInvoiceRejectAPIRequest 机构驳回商家票据信息 API请求
// alibaba.csr.donate.org.invoice.reject
//
// 机构驳回商家票据信息
type AlibabaCsrDonateOrgInvoiceRejectAPIRequest struct {
	model.Params
	// 参数
	_csrInvoiceExternalOrgRejectDto *CsrInvoiceExternalOrgRejectDto
}

// NewAlibabaCsrDonateOrgInvoiceRejectRequest 初始化AlibabaCsrDonateOrgInvoiceRejectAPIRequest对象
func NewAlibabaCsrDonateOrgInvoiceRejectRequest() *AlibabaCsrDonateOrgInvoiceRejectAPIRequest {
	return &AlibabaCsrDonateOrgInvoiceRejectAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCsrDonateOrgInvoiceRejectAPIRequest) Reset() {
	r._csrInvoiceExternalOrgRejectDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCsrDonateOrgInvoiceRejectAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.donate.org.invoice.reject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCsrDonateOrgInvoiceRejectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCsrDonateOrgInvoiceRejectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCsrInvoiceExternalOrgRejectDto is CsrInvoiceExternalOrgRejectDto Setter
// 参数
func (r *AlibabaCsrDonateOrgInvoiceRejectAPIRequest) SetCsrInvoiceExternalOrgRejectDto(_csrInvoiceExternalOrgRejectDto *CsrInvoiceExternalOrgRejectDto) error {
	r._csrInvoiceExternalOrgRejectDto = _csrInvoiceExternalOrgRejectDto
	r.Set("csr_invoice_external_org_reject_dto", _csrInvoiceExternalOrgRejectDto)
	return nil
}

// GetCsrInvoiceExternalOrgRejectDto CsrInvoiceExternalOrgRejectDto Getter
func (r AlibabaCsrDonateOrgInvoiceRejectAPIRequest) GetCsrInvoiceExternalOrgRejectDto() *CsrInvoiceExternalOrgRejectDto {
	return r._csrInvoiceExternalOrgRejectDto
}

var poolAlibabaCsrDonateOrgInvoiceRejectAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCsrDonateOrgInvoiceRejectRequest()
	},
}

// GetAlibabaCsrDonateOrgInvoiceRejectRequest 从 sync.Pool 获取 AlibabaCsrDonateOrgInvoiceRejectAPIRequest
func GetAlibabaCsrDonateOrgInvoiceRejectAPIRequest() *AlibabaCsrDonateOrgInvoiceRejectAPIRequest {
	return poolAlibabaCsrDonateOrgInvoiceRejectAPIRequest.Get().(*AlibabaCsrDonateOrgInvoiceRejectAPIRequest)
}

// ReleaseAlibabaCsrDonateOrgInvoiceRejectAPIRequest 将 AlibabaCsrDonateOrgInvoiceRejectAPIRequest 放入 sync.Pool
func ReleaseAlibabaCsrDonateOrgInvoiceRejectAPIRequest(v *AlibabaCsrDonateOrgInvoiceRejectAPIRequest) {
	v.Reset()
	poolAlibabaCsrDonateOrgInvoiceRejectAPIRequest.Put(v)
}
