package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建税控开通工单 API请求
alibaba.einvoice.flow.tax.create

商户在业务前台订购税控产品后，调用阿里发票此接口，提交税号的入驻开通工单。此接口返回为工单的提交结果，非真正入驻结果。开通结果会在商户完成设备的部署安装 入驻完成后，由阿里发票通过消息异步通知到业务前台。
*/
type AlibabaEinvoiceFlowTaxCreateRequest struct {
    model.Params
    // 工单请求
    _invoiceTaxFlowCreateDto   *InvoiceTaxFlowCreateDTO
}

// 初始化AlibabaEinvoiceFlowTaxCreateRequest对象
func NewAlibabaEinvoiceFlowTaxCreateRequest() *AlibabaEinvoiceFlowTaxCreateRequest{
    return &AlibabaEinvoiceFlowTaxCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceFlowTaxCreateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.flow.tax.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceFlowTaxCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceTaxFlowCreateDto Setter
// 工单请求
func (r *AlibabaEinvoiceFlowTaxCreateRequest) SetInvoiceTaxFlowCreateDto(_invoiceTaxFlowCreateDto *InvoiceTaxFlowCreateDTO) error {
    r._invoiceTaxFlowCreateDto = _invoiceTaxFlowCreateDto
    r.Set("invoice_tax_flow_create_dto", _invoiceTaxFlowCreateDto)
    return nil
}

// InvoiceTaxFlowCreateDto Getter
func (r AlibabaEinvoiceFlowTaxCreateRequest) GetInvoiceTaxFlowCreateDto() *InvoiceTaxFlowCreateDTO {
    return r._invoiceTaxFlowCreateDto
}
