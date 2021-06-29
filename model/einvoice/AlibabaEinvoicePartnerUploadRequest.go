package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商发票上传接口（非授权） API请求
alibaba.einvoice.partner.upload

服务商发票上传接口（非授权）
*/
type AlibabaEinvoicePartnerUploadRequest struct {
    model.Params
    // 原蓝票发票号码
    _normalInvoiceNo   string
    // 原蓝票发票代码
    _normalInvoiceCode   string
    // 销方税号
    _payeeRegisterNo   string
    // 发票数据，upload_type=0且invoiceKind=0电子发票时必填
    _invoiceFileData   *model.File
    // 发票号码，upload_type=0时必填
    _invoiceNo   string
    // 发票代码，upload_type=0时必填
    _invoiceCode   string
    // 开票日期，upload_type=0时必填
    _invoiceDate   string
    // 密码区
    _cipherText   string
    // 机器编号
    _deviceNo   string
    // 校验码
    _antiFakeCode   string
    // 发票类型，upload_type=0且invoiceKind=0电子发票时必填，暂时只支持pdf
    _fileDataType   string
    // 原蓝票的reqIndex
    _reqIndex   string
    // 发票种类，0=电子发票，1=纸质普票，2=纸质专票
    _invoiceKind   int64
    // 上传的类型，0=冲红上传，1=作废上传
    _uploadType   int64
}

// 初始化AlibabaEinvoicePartnerUploadRequest对象
func NewAlibabaEinvoicePartnerUploadRequest() *AlibabaEinvoicePartnerUploadRequest{
    return &AlibabaEinvoicePartnerUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePartnerUploadRequest) GetApiMethodName() string {
    return "alibaba.einvoice.partner.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePartnerUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NormalInvoiceNo Setter
// 原蓝票发票号码
func (r *AlibabaEinvoicePartnerUploadRequest) SetNormalInvoiceNo(_normalInvoiceNo string) error {
    r._normalInvoiceNo = _normalInvoiceNo
    r.Set("normal_invoice_no", _normalInvoiceNo)
    return nil
}

// NormalInvoiceNo Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetNormalInvoiceNo() string {
    return r._normalInvoiceNo
}
// NormalInvoiceCode Setter
// 原蓝票发票代码
func (r *AlibabaEinvoicePartnerUploadRequest) SetNormalInvoiceCode(_normalInvoiceCode string) error {
    r._normalInvoiceCode = _normalInvoiceCode
    r.Set("normal_invoice_code", _normalInvoiceCode)
    return nil
}

// NormalInvoiceCode Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetNormalInvoiceCode() string {
    return r._normalInvoiceCode
}
// PayeeRegisterNo Setter
// 销方税号
func (r *AlibabaEinvoicePartnerUploadRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// InvoiceFileData Setter
// 发票数据，upload_type=0且invoiceKind=0电子发票时必填
func (r *AlibabaEinvoicePartnerUploadRequest) SetInvoiceFileData(_invoiceFileData *model.File) error {
    r._invoiceFileData = _invoiceFileData
    r.Set("invoice_file_data", _invoiceFileData)
    return nil
}

// InvoiceFileData Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetInvoiceFileData() *model.File {
    return r._invoiceFileData
}
// InvoiceNo Setter
// 发票号码，upload_type=0时必填
func (r *AlibabaEinvoicePartnerUploadRequest) SetInvoiceNo(_invoiceNo string) error {
    r._invoiceNo = _invoiceNo
    r.Set("invoice_no", _invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetInvoiceNo() string {
    return r._invoiceNo
}
// InvoiceCode Setter
// 发票代码，upload_type=0时必填
func (r *AlibabaEinvoicePartnerUploadRequest) SetInvoiceCode(_invoiceCode string) error {
    r._invoiceCode = _invoiceCode
    r.Set("invoice_code", _invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetInvoiceCode() string {
    return r._invoiceCode
}
// InvoiceDate Setter
// 开票日期，upload_type=0时必填
func (r *AlibabaEinvoicePartnerUploadRequest) SetInvoiceDate(_invoiceDate string) error {
    r._invoiceDate = _invoiceDate
    r.Set("invoice_date", _invoiceDate)
    return nil
}

// InvoiceDate Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetInvoiceDate() string {
    return r._invoiceDate
}
// CipherText Setter
// 密码区
func (r *AlibabaEinvoicePartnerUploadRequest) SetCipherText(_cipherText string) error {
    r._cipherText = _cipherText
    r.Set("cipher_text", _cipherText)
    return nil
}

// CipherText Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetCipherText() string {
    return r._cipherText
}
// DeviceNo Setter
// 机器编号
func (r *AlibabaEinvoicePartnerUploadRequest) SetDeviceNo(_deviceNo string) error {
    r._deviceNo = _deviceNo
    r.Set("device_no", _deviceNo)
    return nil
}

// DeviceNo Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetDeviceNo() string {
    return r._deviceNo
}
// AntiFakeCode Setter
// 校验码
func (r *AlibabaEinvoicePartnerUploadRequest) SetAntiFakeCode(_antiFakeCode string) error {
    r._antiFakeCode = _antiFakeCode
    r.Set("anti_fake_code", _antiFakeCode)
    return nil
}

// AntiFakeCode Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetAntiFakeCode() string {
    return r._antiFakeCode
}
// FileDataType Setter
// 发票类型，upload_type=0且invoiceKind=0电子发票时必填，暂时只支持pdf
func (r *AlibabaEinvoicePartnerUploadRequest) SetFileDataType(_fileDataType string) error {
    r._fileDataType = _fileDataType
    r.Set("file_data_type", _fileDataType)
    return nil
}

// FileDataType Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetFileDataType() string {
    return r._fileDataType
}
// ReqIndex Setter
// 原蓝票的reqIndex
func (r *AlibabaEinvoicePartnerUploadRequest) SetReqIndex(_reqIndex string) error {
    r._reqIndex = _reqIndex
    r.Set("req_index", _reqIndex)
    return nil
}

// ReqIndex Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetReqIndex() string {
    return r._reqIndex
}
// InvoiceKind Setter
// 发票种类，0=电子发票，1=纸质普票，2=纸质专票
func (r *AlibabaEinvoicePartnerUploadRequest) SetInvoiceKind(_invoiceKind int64) error {
    r._invoiceKind = _invoiceKind
    r.Set("invoice_kind", _invoiceKind)
    return nil
}

// InvoiceKind Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetInvoiceKind() int64 {
    return r._invoiceKind
}
// UploadType Setter
// 上传的类型，0=冲红上传，1=作废上传
func (r *AlibabaEinvoicePartnerUploadRequest) SetUploadType(_uploadType int64) error {
    r._uploadType = _uploadType
    r.Set("upload_type", _uploadType)
    return nil
}

// UploadType Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetUploadType() int64 {
    return r._uploadType
}
