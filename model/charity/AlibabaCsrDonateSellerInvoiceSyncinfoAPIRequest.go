package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacsrdonatesellerinvoicesyncinfoAPIRequest 链上同步商家票据信息 API请求
// alibaba.csr.donate.seller.invoice.syncinfo
//
// 链上同步商家票据信息
type AlibabacsrdonatesellerinvoicesyncinfoAPIRequest struct {
	model.Params
	// 请求request
	_csrInvoiceAntChainSyncDto *CsrInvoiceAntChainSyncDto
}

// NewAlibabacsrdonatesellerinvoicesyncinfoRequest 初始化AlibabacsrdonatesellerinvoicesyncinfoAPIRequest对象
func NewAlibabacsrdonatesellerinvoicesyncinfoRequest() *AlibabacsrdonatesellerinvoicesyncinfoAPIRequest {
	return &AlibabacsrdonatesellerinvoicesyncinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacsrdonatesellerinvoicesyncinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.donate.seller.invoice.syncinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacsrdonatesellerinvoicesyncinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacsrdonatesellerinvoicesyncinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCsrInvoiceAntChainSyncDto is CsrInvoiceAntChainSyncDto Setter
// 请求request
func (r *AlibabacsrdonatesellerinvoicesyncinfoAPIRequest) SetCsrInvoiceAntChainSyncDto(_csrInvoiceAntChainSyncDto *CsrInvoiceAntChainSyncDto) error {
	r._csrInvoiceAntChainSyncDto = _csrInvoiceAntChainSyncDto
	r.Set("csr_invoice_ant_chain_sync_dto", _csrInvoiceAntChainSyncDto)
	return nil
}

// GetCsrInvoiceAntChainSyncDto CsrInvoiceAntChainSyncDto Getter
func (r AlibabacsrdonatesellerinvoicesyncinfoAPIRequest) GetCsrInvoiceAntChainSyncDto() *CsrInvoiceAntChainSyncDto {
	return r._csrInvoiceAntChainSyncDto
}
