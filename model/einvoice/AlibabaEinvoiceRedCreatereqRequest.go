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
type AlibabaEinvoiceRedCreatereqRequest struct {
    model.Params
    // 销售方税号
    payeeRegisterNo   string
    // 蓝票流水号，优先级高于发票代码+发票号码
    blueSerialNo   string
    // 红票流水号
    redSerialNo   string
    // 蓝票发票代码
    invoiceCode   string
    // 蓝票发票号码
    invoiceNo   string
}

// 初始化AlibabaEinvoiceRedCreatereqRequest对象
func NewAlibabaEinvoiceRedCreatereqRequest() *AlibabaEinvoiceRedCreatereqRequest{
    return &AlibabaEinvoiceRedCreatereqRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceRedCreatereqRequest) GetApiMethodName() string {
    return "alibaba.einvoice.red.createreq"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceRedCreatereqRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PayeeRegisterNo Setter
// 销售方税号
func (r *AlibabaEinvoiceRedCreatereqRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceRedCreatereqRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// BlueSerialNo Setter
// 蓝票流水号，优先级高于发票代码+发票号码
func (r *AlibabaEinvoiceRedCreatereqRequest) SetBlueSerialNo(blueSerialNo string) error {
    r.blueSerialNo = blueSerialNo
    r.Set("blue_serial_no", blueSerialNo)
    return nil
}

// BlueSerialNo Getter
func (r AlibabaEinvoiceRedCreatereqRequest) GetBlueSerialNo() string {
    return r.blueSerialNo
}
// RedSerialNo Setter
// 红票流水号
func (r *AlibabaEinvoiceRedCreatereqRequest) SetRedSerialNo(redSerialNo string) error {
    r.redSerialNo = redSerialNo
    r.Set("red_serial_no", redSerialNo)
    return nil
}

// RedSerialNo Getter
func (r AlibabaEinvoiceRedCreatereqRequest) GetRedSerialNo() string {
    return r.redSerialNo
}
// InvoiceCode Setter
// 蓝票发票代码
func (r *AlibabaEinvoiceRedCreatereqRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoiceRedCreatereqRequest) GetInvoiceCode() string {
    return r.invoiceCode
}
// InvoiceNo Setter
// 蓝票发票号码
func (r *AlibabaEinvoiceRedCreatereqRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoiceRedCreatereqRequest) GetInvoiceNo() string {
    return r.invoiceNo
}
