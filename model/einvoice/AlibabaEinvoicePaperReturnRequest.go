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
type AlibabaEinvoicePaperReturnAPIRequest struct {
    model.Params
    // 发票密文，密码区的字符串
    _ciphertext   string
    // 发票号码
    _invoiceNo   string
    // 发票日期
    _invoiceDate   string
    // 防伪码
    _antiFakeCode   string
    // 税控设备编号(新版电子发票有)
    _deviceNo   string
    // 发票代码
    _invoiceCode   string
    // 开票结果"success"或者"fail"
    _createResult   string
    // 错误码
    _bizErrorCode   string
    // 错误信息
    _bizErrorMsg   string
    // 开票请求的唯一索引
    _reqIndex   string
}

// 初始化AlibabaEinvoicePaperReturnAPIRequest对象
func NewAlibabaEinvoicePaperReturnRequest() *AlibabaEinvoicePaperReturnAPIRequest{
    return &AlibabaEinvoicePaperReturnAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePaperReturnAPIRequest) GetApiMethodName() string {
    return "alibaba.einvoice.paper.return"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePaperReturnAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ciphertext Setter
// 发票密文，密码区的字符串
func (r *AlibabaEinvoicePaperReturnAPIRequest) SetCiphertext(_ciphertext string) error {
    r._ciphertext = _ciphertext
    r.Set("ciphertext", _ciphertext)
    return nil
}

// Ciphertext Getter
func (r AlibabaEinvoicePaperReturnAPIRequest) GetCiphertext() string {
    return r._ciphertext
}
// InvoiceNo Setter
// 发票号码
func (r *AlibabaEinvoicePaperReturnAPIRequest) SetInvoiceNo(_invoiceNo string) error {
    r._invoiceNo = _invoiceNo
    r.Set("invoice_no", _invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoicePaperReturnAPIRequest) GetInvoiceNo() string {
    return r._invoiceNo
}
// InvoiceDate Setter
// 发票日期
func (r *AlibabaEinvoicePaperReturnAPIRequest) SetInvoiceDate(_invoiceDate string) error {
    r._invoiceDate = _invoiceDate
    r.Set("invoice_date", _invoiceDate)
    return nil
}

// InvoiceDate Getter
func (r AlibabaEinvoicePaperReturnAPIRequest) GetInvoiceDate() string {
    return r._invoiceDate
}
// AntiFakeCode Setter
// 防伪码
func (r *AlibabaEinvoicePaperReturnAPIRequest) SetAntiFakeCode(_antiFakeCode string) error {
    r._antiFakeCode = _antiFakeCode
    r.Set("anti_fake_code", _antiFakeCode)
    return nil
}

// AntiFakeCode Getter
func (r AlibabaEinvoicePaperReturnAPIRequest) GetAntiFakeCode() string {
    return r._antiFakeCode
}
// DeviceNo Setter
// 税控设备编号(新版电子发票有)
func (r *AlibabaEinvoicePaperReturnAPIRequest) SetDeviceNo(_deviceNo string) error {
    r._deviceNo = _deviceNo
    r.Set("device_no", _deviceNo)
    return nil
}

// DeviceNo Getter
func (r AlibabaEinvoicePaperReturnAPIRequest) GetDeviceNo() string {
    return r._deviceNo
}
// InvoiceCode Setter
// 发票代码
func (r *AlibabaEinvoicePaperReturnAPIRequest) SetInvoiceCode(_invoiceCode string) error {
    r._invoiceCode = _invoiceCode
    r.Set("invoice_code", _invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoicePaperReturnAPIRequest) GetInvoiceCode() string {
    return r._invoiceCode
}
// CreateResult Setter
// 开票结果"success"或者"fail"
func (r *AlibabaEinvoicePaperReturnAPIRequest) SetCreateResult(_createResult string) error {
    r._createResult = _createResult
    r.Set("create_result", _createResult)
    return nil
}

// CreateResult Getter
func (r AlibabaEinvoicePaperReturnAPIRequest) GetCreateResult() string {
    return r._createResult
}
// BizErrorCode Setter
// 错误码
func (r *AlibabaEinvoicePaperReturnAPIRequest) SetBizErrorCode(_bizErrorCode string) error {
    r._bizErrorCode = _bizErrorCode
    r.Set("biz_error_code", _bizErrorCode)
    return nil
}

// BizErrorCode Getter
func (r AlibabaEinvoicePaperReturnAPIRequest) GetBizErrorCode() string {
    return r._bizErrorCode
}
// BizErrorMsg Setter
// 错误信息
func (r *AlibabaEinvoicePaperReturnAPIRequest) SetBizErrorMsg(_bizErrorMsg string) error {
    r._bizErrorMsg = _bizErrorMsg
    r.Set("biz_error_msg", _bizErrorMsg)
    return nil
}

// BizErrorMsg Getter
func (r AlibabaEinvoicePaperReturnAPIRequest) GetBizErrorMsg() string {
    return r._bizErrorMsg
}
// ReqIndex Setter
// 开票请求的唯一索引
func (r *AlibabaEinvoicePaperReturnAPIRequest) SetReqIndex(_reqIndex string) error {
    r._reqIndex = _reqIndex
    r.Set("req_index", _reqIndex)
    return nil
}

// ReqIndex Getter
func (r AlibabaEinvoicePaperReturnAPIRequest) GetReqIndex() string {
    return r._reqIndex
}
