package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
纸质发票结果回传 APIRequest
alibaba.einvoice.paper.return

纸质发票结果回传
*/
type AlibabaEinvoicePaperReturnRequest struct {
    model.Params

    // 发票密文，密码区的字符串
    ciphertext   string 

    // 发票号码
    invoiceNo   string 

    // 发票日期
    invoiceDate   string 

    // 防伪码
    antiFakeCode   string 

    // 税控设备编号(新版电子发票有)
    deviceNo   string 

    // 发票代码
    invoiceCode   string 

    // 开票结果"success"或者"fail"
    createResult   string 

    // 错误码
    bizErrorCode   string 

    // 错误信息
    bizErrorMsg   string 

    // 开票请求的唯一索引
    reqIndex   string 

}

func NewAlibabaEinvoicePaperReturnRequest() *AlibabaEinvoicePaperReturnRequest{
    return &AlibabaEinvoicePaperReturnRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoicePaperReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.paper.return"
}

func (r AlibabaEinvoicePaperReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoicePaperReturnRequest) SetCiphertext(ciphertext string) error {
    r.ciphertext = ciphertext
    r.Set("ciphertext", ciphertext)
    return nil
}

func (r AlibabaEinvoicePaperReturnRequest) GetCiphertext() string {
    return r.ciphertext
}

func (r *AlibabaEinvoicePaperReturnRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

func (r AlibabaEinvoicePaperReturnRequest) GetInvoiceNo() string {
    return r.invoiceNo
}

func (r *AlibabaEinvoicePaperReturnRequest) SetInvoiceDate(invoiceDate string) error {
    r.invoiceDate = invoiceDate
    r.Set("invoice_date", invoiceDate)
    return nil
}

func (r AlibabaEinvoicePaperReturnRequest) GetInvoiceDate() string {
    return r.invoiceDate
}

func (r *AlibabaEinvoicePaperReturnRequest) SetAntiFakeCode(antiFakeCode string) error {
    r.antiFakeCode = antiFakeCode
    r.Set("anti_fake_code", antiFakeCode)
    return nil
}

func (r AlibabaEinvoicePaperReturnRequest) GetAntiFakeCode() string {
    return r.antiFakeCode
}

func (r *AlibabaEinvoicePaperReturnRequest) SetDeviceNo(deviceNo string) error {
    r.deviceNo = deviceNo
    r.Set("device_no", deviceNo)
    return nil
}

func (r AlibabaEinvoicePaperReturnRequest) GetDeviceNo() string {
    return r.deviceNo
}

func (r *AlibabaEinvoicePaperReturnRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

func (r AlibabaEinvoicePaperReturnRequest) GetInvoiceCode() string {
    return r.invoiceCode
}

func (r *AlibabaEinvoicePaperReturnRequest) SetCreateResult(createResult string) error {
    r.createResult = createResult
    r.Set("create_result", createResult)
    return nil
}

func (r AlibabaEinvoicePaperReturnRequest) GetCreateResult() string {
    return r.createResult
}

func (r *AlibabaEinvoicePaperReturnRequest) SetBizErrorCode(bizErrorCode string) error {
    r.bizErrorCode = bizErrorCode
    r.Set("biz_error_code", bizErrorCode)
    return nil
}

func (r AlibabaEinvoicePaperReturnRequest) GetBizErrorCode() string {
    return r.bizErrorCode
}

func (r *AlibabaEinvoicePaperReturnRequest) SetBizErrorMsg(bizErrorMsg string) error {
    r.bizErrorMsg = bizErrorMsg
    r.Set("biz_error_msg", bizErrorMsg)
    return nil
}

func (r AlibabaEinvoicePaperReturnRequest) GetBizErrorMsg() string {
    return r.bizErrorMsg
}

func (r *AlibabaEinvoicePaperReturnRequest) SetReqIndex(reqIndex string) error {
    r.reqIndex = reqIndex
    r.Set("req_index", reqIndex)
    return nil
}

func (r AlibabaEinvoicePaperReturnRequest) GetReqIndex() string {
    return r.reqIndex
}

