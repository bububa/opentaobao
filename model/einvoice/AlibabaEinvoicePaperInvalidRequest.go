package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
纸票作废接口 APIRequest
alibaba.einvoice.paper.invalid

作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票
*/
type AlibabaEinvoicePaperInvalidRequest struct {
    model.Params

    // 发票代码，空白作废时必填
    invoiceCode   string 

    // 发票号码，空白作废时必填
    invoiceNo   string 

    // 作废操作人
    invalidOperator   string 

    // 作废类型, 0=空白发票(有残缺 的纸张发票，不能做为有效报销)作废, 1=已开发票作废
    invalidType   int64 

    // 销售方纳税人识别号
    payeeRegisterNo   string 

    // 开票流水号
    serialNo   string 

}

func NewAlibabaEinvoicePaperInvalidRequest() *AlibabaEinvoicePaperInvalidRequest{
    return &AlibabaEinvoicePaperInvalidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoicePaperInvalidRequest) GetApiMethodName() string {
    return "alibaba.einvoice.paper.invalid"
}

func (r AlibabaEinvoicePaperInvalidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoicePaperInvalidRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

func (r AlibabaEinvoicePaperInvalidRequest) GetInvoiceCode() string {
    return r.invoiceCode
}

func (r *AlibabaEinvoicePaperInvalidRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

func (r AlibabaEinvoicePaperInvalidRequest) GetInvoiceNo() string {
    return r.invoiceNo
}

func (r *AlibabaEinvoicePaperInvalidRequest) SetInvalidOperator(invalidOperator string) error {
    r.invalidOperator = invalidOperator
    r.Set("invalid_operator", invalidOperator)
    return nil
}

func (r AlibabaEinvoicePaperInvalidRequest) GetInvalidOperator() string {
    return r.invalidOperator
}

func (r *AlibabaEinvoicePaperInvalidRequest) SetInvalidType(invalidType int64) error {
    r.invalidType = invalidType
    r.Set("invalid_type", invalidType)
    return nil
}

func (r AlibabaEinvoicePaperInvalidRequest) GetInvalidType() int64 {
    return r.invalidType
}

func (r *AlibabaEinvoicePaperInvalidRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoicePaperInvalidRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoicePaperInvalidRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

func (r AlibabaEinvoicePaperInvalidRequest) GetSerialNo() string {
    return r.serialNo
}

