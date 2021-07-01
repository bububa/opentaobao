package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票冲红接口 API请求
alibaba.einvoice.red.createreq

发票冲红接口，通过蓝票流水号或者发票号码+发票代码进行冲红
*/
type AlibabaEinvoiceRedCreatereqAPIRequest struct {
    model.Params
    // 销售方税号
    _payeeRegisterNo   string
    // 蓝票流水号，优先级高于发票代码+发票号码
    _blueSerialNo   string
    // 红票流水号
    _redSerialNo   string
    // 蓝票发票代码
    _invoiceCode   string
    // 蓝票发票号码
    _invoiceNo   string
}

// 初始化AlibabaEinvoiceRedCreatereqAPIRequest对象
func NewAlibabaEinvoiceRedCreatereqRequest() *AlibabaEinvoiceRedCreatereqAPIRequest{
    return &AlibabaEinvoiceRedCreatereqAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetApiMethodName() string {
    return "alibaba.einvoice.red.createreq"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PayeeRegisterNo Setter
// 销售方税号
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// BlueSerialNo Setter
// 蓝票流水号，优先级高于发票代码+发票号码
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) SetBlueSerialNo(_blueSerialNo string) error {
    r._blueSerialNo = _blueSerialNo
    r.Set("blue_serial_no", _blueSerialNo)
    return nil
}

// BlueSerialNo Getter
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetBlueSerialNo() string {
    return r._blueSerialNo
}
// RedSerialNo Setter
// 红票流水号
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) SetRedSerialNo(_redSerialNo string) error {
    r._redSerialNo = _redSerialNo
    r.Set("red_serial_no", _redSerialNo)
    return nil
}

// RedSerialNo Getter
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetRedSerialNo() string {
    return r._redSerialNo
}
// InvoiceCode Setter
// 蓝票发票代码
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) SetInvoiceCode(_invoiceCode string) error {
    r._invoiceCode = _invoiceCode
    r.Set("invoice_code", _invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetInvoiceCode() string {
    return r._invoiceCode
}
// InvoiceNo Setter
// 蓝票发票号码
func (r *AlibabaEinvoiceRedCreatereqAPIRequest) SetInvoiceNo(_invoiceNo string) error {
    r._invoiceNo = _invoiceNo
    r.Set("invoice_no", _invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoiceRedCreatereqAPIRequest) GetInvoiceNo() string {
    return r._invoiceNo
}
