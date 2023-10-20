package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceFlowTaxCreateAPIRequest 创建税控开通工单 API请求
// alibaba.einvoice.flow.tax.create
//
// 商户在业务前台订购税控产品后，调用阿里发票此接口，提交税号的入驻开通工单。此接口返回为工单的提交结果，非真正入驻结果。开通结果会在商户完成设备的部署安装 入驻完成后，由阿里发票通过消息异步通知到业务前台。
type AlibabaEinvoiceFlowTaxCreateAPIRequest struct {
	model.Params
	// 工单请求
	_invoiceTaxFlowCreateDto *InvoiceTaxFlowCreateDto
}

// NewAlibabaEinvoiceFlowTaxCreateRequest 初始化AlibabaEinvoiceFlowTaxCreateAPIRequest对象
func NewAlibabaEinvoiceFlowTaxCreateRequest() *AlibabaEinvoiceFlowTaxCreateAPIRequest {
	return &AlibabaEinvoiceFlowTaxCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceFlowTaxCreateAPIRequest) Reset() {
	r._invoiceTaxFlowCreateDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceFlowTaxCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.flow.tax.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceFlowTaxCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceFlowTaxCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceTaxFlowCreateDto is InvoiceTaxFlowCreateDto Setter
// 工单请求
func (r *AlibabaEinvoiceFlowTaxCreateAPIRequest) SetInvoiceTaxFlowCreateDto(_invoiceTaxFlowCreateDto *InvoiceTaxFlowCreateDto) error {
	r._invoiceTaxFlowCreateDto = _invoiceTaxFlowCreateDto
	r.Set("invoice_tax_flow_create_dto", _invoiceTaxFlowCreateDto)
	return nil
}

// GetInvoiceTaxFlowCreateDto InvoiceTaxFlowCreateDto Getter
func (r AlibabaEinvoiceFlowTaxCreateAPIRequest) GetInvoiceTaxFlowCreateDto() *InvoiceTaxFlowCreateDto {
	return r._invoiceTaxFlowCreateDto
}

var poolAlibabaEinvoiceFlowTaxCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceFlowTaxCreateRequest()
	},
}

// GetAlibabaEinvoiceFlowTaxCreateRequest 从 sync.Pool 获取 AlibabaEinvoiceFlowTaxCreateAPIRequest
func GetAlibabaEinvoiceFlowTaxCreateAPIRequest() *AlibabaEinvoiceFlowTaxCreateAPIRequest {
	return poolAlibabaEinvoiceFlowTaxCreateAPIRequest.Get().(*AlibabaEinvoiceFlowTaxCreateAPIRequest)
}

// ReleaseAlibabaEinvoiceFlowTaxCreateAPIRequest 将 AlibabaEinvoiceFlowTaxCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceFlowTaxCreateAPIRequest(v *AlibabaEinvoiceFlowTaxCreateAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceFlowTaxCreateAPIRequest.Put(v)
}
