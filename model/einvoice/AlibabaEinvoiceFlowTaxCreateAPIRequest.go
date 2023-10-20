package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceflowtaxcreateAPIRequest 创建税控开通工单 API请求
// alibaba.einvoice.flow.tax.create
//
// 商户在业务前台订购税控产品后，调用阿里发票此接口，提交税号的入驻开通工单。此接口返回为工单的提交结果，非真正入驻结果。开通结果会在商户完成设备的部署安装 入驻完成后，由阿里发票通过消息异步通知到业务前台。
type AlibabaeinvoiceflowtaxcreateAPIRequest struct {
	model.Params
	// 工单请求
	_invoiceTaxFlowCreateDto *InvoiceTaxFlowCreateDto
}

// NewAlibabaeinvoiceflowtaxcreateRequest 初始化AlibabaeinvoiceflowtaxcreateAPIRequest对象
func NewAlibabaeinvoiceflowtaxcreateRequest() *AlibabaeinvoiceflowtaxcreateAPIRequest {
	return &AlibabaeinvoiceflowtaxcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceflowtaxcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.flow.tax.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceflowtaxcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceflowtaxcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceTaxFlowCreateDto is InvoiceTaxFlowCreateDto Setter
// 工单请求
func (r *AlibabaeinvoiceflowtaxcreateAPIRequest) SetInvoiceTaxFlowCreateDto(_invoiceTaxFlowCreateDto *InvoiceTaxFlowCreateDto) error {
	r._invoiceTaxFlowCreateDto = _invoiceTaxFlowCreateDto
	r.Set("invoice_tax_flow_create_dto", _invoiceTaxFlowCreateDto)
	return nil
}

// GetInvoiceTaxFlowCreateDto InvoiceTaxFlowCreateDto Getter
func (r AlibabaeinvoiceflowtaxcreateAPIRequest) GetInvoiceTaxFlowCreateDto() *InvoiceTaxFlowCreateDto {
	return r._invoiceTaxFlowCreateDto
}
