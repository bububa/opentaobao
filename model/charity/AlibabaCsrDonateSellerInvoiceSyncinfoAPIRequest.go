package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest 链上同步商家票据信息 API请求
// alibaba.csr.donate.seller.invoice.syncinfo
//
// 链上同步商家票据信息
type AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest struct {
	model.Params
	// 请求request
	_csrInvoiceAntChainSyncDto *CsrInvoiceAntChainSyncDto
}

// NewAlibabaCsrDonateSellerInvoiceSyncinfoRequest 初始化AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest对象
func NewAlibabaCsrDonateSellerInvoiceSyncinfoRequest() *AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest {
	return &AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.donate.seller.invoice.syncinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCsrInvoiceAntChainSyncDto is CsrInvoiceAntChainSyncDto Setter
// 请求request
func (r *AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest) SetCsrInvoiceAntChainSyncDto(_csrInvoiceAntChainSyncDto *CsrInvoiceAntChainSyncDto) error {
	r._csrInvoiceAntChainSyncDto = _csrInvoiceAntChainSyncDto
	r.Set("csr_invoice_ant_chain_sync_dto", _csrInvoiceAntChainSyncDto)
	return nil
}

// GetCsrInvoiceAntChainSyncDto CsrInvoiceAntChainSyncDto Getter
func (r AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest) GetCsrInvoiceAntChainSyncDto() *CsrInvoiceAntChainSyncDto {
	return r._csrInvoiceAntChainSyncDto
}
