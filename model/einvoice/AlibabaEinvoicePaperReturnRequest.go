package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
纸质发票结果回传 API请求
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

// 初始化AlibabaEinvoicePaperReturnRequest对象
func NewAlibabaEinvoicePaperReturnRequest() *AlibabaEinvoicePaperReturnRequest{
    return &AlibabaEinvoicePaperReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePaperReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.paper.return"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePaperReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ciphertext Setter
// 发票密文，密码区的字符串
func (r *AlibabaEinvoicePaperReturnRequest) SetCiphertext(ciphertext string) error {
    r.ciphertext = ciphertext
    r.Set("ciphertext", ciphertext)
    return nil
}

// Ciphertext Getter
func (r AlibabaEinvoicePaperReturnRequest) GetCiphertext() string {
    return r.ciphertext
}
// InvoiceNo Setter
// 发票号码
func (r *AlibabaEinvoicePaperReturnRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoicePaperReturnRequest) GetInvoiceNo() string {
    return r.invoiceNo
}
// InvoiceDate Setter
// 发票日期
func (r *AlibabaEinvoicePaperReturnRequest) SetInvoiceDate(invoiceDate string) error {
    r.invoiceDate = invoiceDate
    r.Set("invoice_date", invoiceDate)
    return nil
}

// InvoiceDate Getter
func (r AlibabaEinvoicePaperReturnRequest) GetInvoiceDate() string {
    return r.invoiceDate
}
// AntiFakeCode Setter
// 防伪码
func (r *AlibabaEinvoicePaperReturnRequest) SetAntiFakeCode(antiFakeCode string) error {
    r.antiFakeCode = antiFakeCode
    r.Set("anti_fake_code", antiFakeCode)
    return nil
}

// AntiFakeCode Getter
func (r AlibabaEinvoicePaperReturnRequest) GetAntiFakeCode() string {
    return r.antiFakeCode
}
// DeviceNo Setter
// 税控设备编号(新版电子发票有)
func (r *AlibabaEinvoicePaperReturnRequest) SetDeviceNo(deviceNo string) error {
    r.deviceNo = deviceNo
    r.Set("device_no", deviceNo)
    return nil
}

// DeviceNo Getter
func (r AlibabaEinvoicePaperReturnRequest) GetDeviceNo() string {
    return r.deviceNo
}
// InvoiceCode Setter
// 发票代码
func (r *AlibabaEinvoicePaperReturnRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoicePaperReturnRequest) GetInvoiceCode() string {
    return r.invoiceCode
}
// CreateResult Setter
// 开票结果"success"或者"fail"
func (r *AlibabaEinvoicePaperReturnRequest) SetCreateResult(createResult string) error {
    r.createResult = createResult
    r.Set("create_result", createResult)
    return nil
}

// CreateResult Getter
func (r AlibabaEinvoicePaperReturnRequest) GetCreateResult() string {
    return r.createResult
}
// BizErrorCode Setter
// 错误码
func (r *AlibabaEinvoicePaperReturnRequest) SetBizErrorCode(bizErrorCode string) error {
    r.bizErrorCode = bizErrorCode
    r.Set("biz_error_code", bizErrorCode)
    return nil
}

// BizErrorCode Getter
func (r AlibabaEinvoicePaperReturnRequest) GetBizErrorCode() string {
    return r.bizErrorCode
}
// BizErrorMsg Setter
// 错误信息
func (r *AlibabaEinvoicePaperReturnRequest) SetBizErrorMsg(bizErrorMsg string) error {
    r.bizErrorMsg = bizErrorMsg
    r.Set("biz_error_msg", bizErrorMsg)
    return nil
}

// BizErrorMsg Getter
func (r AlibabaEinvoicePaperReturnRequest) GetBizErrorMsg() string {
    return r.bizErrorMsg
}
// ReqIndex Setter
// 开票请求的唯一索引
func (r *AlibabaEinvoicePaperReturnRequest) SetReqIndex(reqIndex string) error {
    r.reqIndex = reqIndex
    r.Set("req_index", reqIndex)
    return nil
}

// ReqIndex Getter
func (r AlibabaEinvoicePaperReturnRequest) GetReqIndex() string {
    return r.reqIndex
}
