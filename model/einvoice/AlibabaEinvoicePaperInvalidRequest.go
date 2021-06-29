package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
纸票作废接口 API请求
alibaba.einvoice.paper.invalid

作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票
*/
type AlibabaEinvoicePaperInvalidRequest struct {
    model.Params
    // 发票代码，空白作废时必填
    _invoiceCode   string
    // 发票号码，空白作废时必填
    _invoiceNo   string
    // 作废操作人
    _invalidOperator   string
    // 作废类型, 0=空白发票(有残缺 的纸张发票，不能做为有效报销)作废, 1=已开发票作废
    _invalidType   int64
    // 销售方纳税人识别号
    _payeeRegisterNo   string
    // 开票流水号
    _serialNo   string
}

// 初始化AlibabaEinvoicePaperInvalidRequest对象
func NewAlibabaEinvoicePaperInvalidRequest() *AlibabaEinvoicePaperInvalidRequest{
    return &AlibabaEinvoicePaperInvalidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePaperInvalidRequest) GetApiMethodName() string {
    return "alibaba.einvoice.paper.invalid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePaperInvalidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceCode Setter
// 发票代码，空白作废时必填
func (r *AlibabaEinvoicePaperInvalidRequest) SetInvoiceCode(_invoiceCode string) error {
    r._invoiceCode = _invoiceCode
    r.Set("invoice_code", _invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoicePaperInvalidRequest) GetInvoiceCode() string {
    return r._invoiceCode
}
// InvoiceNo Setter
// 发票号码，空白作废时必填
func (r *AlibabaEinvoicePaperInvalidRequest) SetInvoiceNo(_invoiceNo string) error {
    r._invoiceNo = _invoiceNo
    r.Set("invoice_no", _invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoicePaperInvalidRequest) GetInvoiceNo() string {
    return r._invoiceNo
}
// InvalidOperator Setter
// 作废操作人
func (r *AlibabaEinvoicePaperInvalidRequest) SetInvalidOperator(_invalidOperator string) error {
    r._invalidOperator = _invalidOperator
    r.Set("invalid_operator", _invalidOperator)
    return nil
}

// InvalidOperator Getter
func (r AlibabaEinvoicePaperInvalidRequest) GetInvalidOperator() string {
    return r._invalidOperator
}
// InvalidType Setter
// 作废类型, 0=空白发票(有残缺 的纸张发票，不能做为有效报销)作废, 1=已开发票作废
func (r *AlibabaEinvoicePaperInvalidRequest) SetInvalidType(_invalidType int64) error {
    r._invalidType = _invalidType
    r.Set("invalid_type", _invalidType)
    return nil
}

// InvalidType Getter
func (r AlibabaEinvoicePaperInvalidRequest) GetInvalidType() int64 {
    return r._invalidType
}
// PayeeRegisterNo Setter
// 销售方纳税人识别号
func (r *AlibabaEinvoicePaperInvalidRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoicePaperInvalidRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// SerialNo Setter
// 开票流水号
func (r *AlibabaEinvoicePaperInvalidRequest) SetSerialNo(_serialNo string) error {
    r._serialNo = _serialNo
    r.Set("serial_no", _serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoicePaperInvalidRequest) GetSerialNo() string {
    return r._serialNo
}
