package charity

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest 获取机构待开票列表 API请求
// alibaba.csr.donate.org.invoice.undraw.list
//
// 获取机构待开票列表
type AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest struct {
	model.Params
	// 请求参数
	_csrInvoiceExternalOrgQueryDto *CsrInvoiceExternalOrgQueryDto
}

// NewAlibabaCsrDonateOrgInvoiceUndrawListRequest 初始化AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest对象
func NewAlibabaCsrDonateOrgInvoiceUndrawListRequest() *AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest {
	return &AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest) Reset() {
	r._csrInvoiceExternalOrgQueryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.donate.org.invoice.undraw.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCsrInvoiceExternalOrgQueryDto is CsrInvoiceExternalOrgQueryDto Setter
// 请求参数
func (r *AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest) SetCsrInvoiceExternalOrgQueryDto(_csrInvoiceExternalOrgQueryDto *CsrInvoiceExternalOrgQueryDto) error {
	r._csrInvoiceExternalOrgQueryDto = _csrInvoiceExternalOrgQueryDto
	r.Set("csr_invoice_external_org_query_dto", _csrInvoiceExternalOrgQueryDto)
	return nil
}

// GetCsrInvoiceExternalOrgQueryDto CsrInvoiceExternalOrgQueryDto Getter
func (r AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest) GetCsrInvoiceExternalOrgQueryDto() *CsrInvoiceExternalOrgQueryDto {
	return r._csrInvoiceExternalOrgQueryDto
}

var poolAlibabaCsrDonateOrgInvoiceUndrawListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCsrDonateOrgInvoiceUndrawListRequest()
	},
}

// GetAlibabaCsrDonateOrgInvoiceUndrawListRequest 从 sync.Pool 获取 AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest
func GetAlibabaCsrDonateOrgInvoiceUndrawListAPIRequest() *AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest {
	return poolAlibabaCsrDonateOrgInvoiceUndrawListAPIRequest.Get().(*AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest)
}

// ReleaseAlibabaCsrDonateOrgInvoiceUndrawListAPIRequest 将 AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest 放入 sync.Pool
func ReleaseAlibabaCsrDonateOrgInvoiceUndrawListAPIRequest(v *AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest) {
	v.Reset()
	poolAlibabaCsrDonateOrgInvoiceUndrawListAPIRequest.Put(v)
}
