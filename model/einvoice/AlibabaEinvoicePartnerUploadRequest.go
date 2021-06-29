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
    normalInvoiceNo   string
    // 原蓝票发票代码
    normalInvoiceCode   string
    // 销方税号
    payeeRegisterNo   string
    // 发票数据，upload_type=0且invoiceKind=0电子发票时必填
    invoiceFileData   []*model.File
    // 发票号码，upload_type=0时必填
    invoiceNo   string
    // 发票代码，upload_type=0时必填
    invoiceCode   string
    // 开票日期，upload_type=0时必填
    invoiceDate   string
    // 密码区
    cipherText   string
    // 机器编号
    deviceNo   string
    // 校验码
    antiFakeCode   string
    // 发票类型，upload_type=0且invoiceKind=0电子发票时必填，暂时只支持pdf
    fileDataType   string
    // 原蓝票的reqIndex
    reqIndex   string
    // 发票种类，0=电子发票，1=纸质普票，2=纸质专票
    invoiceKind   int64
    // 上传的类型，0=冲红上传，1=作废上传
    uploadType   int64
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
func (r *AlibabaEinvoicePartnerUploadRequest) SetNormalInvoiceNo(normalInvoiceNo string) error {
    r.normalInvoiceNo = normalInvoiceNo
    r.Set("normal_invoice_no", normalInvoiceNo)
    return nil
}

// NormalInvoiceNo Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetNormalInvoiceNo() string {
    return r.normalInvoiceNo
}
// NormalInvoiceCode Setter
// 原蓝票发票代码
func (r *AlibabaEinvoicePartnerUploadRequest) SetNormalInvoiceCode(normalInvoiceCode string) error {
    r.normalInvoiceCode = normalInvoiceCode
    r.Set("normal_invoice_code", normalInvoiceCode)
    return nil
}

// NormalInvoiceCode Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetNormalInvoiceCode() string {
    return r.normalInvoiceCode
}
// PayeeRegisterNo Setter
// 销方税号
func (r *AlibabaEinvoicePartnerUploadRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// InvoiceFileData Setter
// 发票数据，upload_type=0且invoiceKind=0电子发票时必填
func (r *AlibabaEinvoicePartnerUploadRequest) SetInvoiceFileData(invoiceFileData []*model.File) error {
    r.invoiceFileData = invoiceFileData
    r.Set("invoice_file_data", invoiceFileData)
    return nil
}

// InvoiceFileData Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetInvoiceFileData() []*model.File {
    return r.invoiceFileData
}
// InvoiceNo Setter
// 发票号码，upload_type=0时必填
func (r *AlibabaEinvoicePartnerUploadRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetInvoiceNo() string {
    return r.invoiceNo
}
// InvoiceCode Setter
// 发票代码，upload_type=0时必填
func (r *AlibabaEinvoicePartnerUploadRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetInvoiceCode() string {
    return r.invoiceCode
}
// InvoiceDate Setter
// 开票日期，upload_type=0时必填
func (r *AlibabaEinvoicePartnerUploadRequest) SetInvoiceDate(invoiceDate string) error {
    r.invoiceDate = invoiceDate
    r.Set("invoice_date", invoiceDate)
    return nil
}

// InvoiceDate Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetInvoiceDate() string {
    return r.invoiceDate
}
// CipherText Setter
// 密码区
func (r *AlibabaEinvoicePartnerUploadRequest) SetCipherText(cipherText string) error {
    r.cipherText = cipherText
    r.Set("cipher_text", cipherText)
    return nil
}

// CipherText Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetCipherText() string {
    return r.cipherText
}
// DeviceNo Setter
// 机器编号
func (r *AlibabaEinvoicePartnerUploadRequest) SetDeviceNo(deviceNo string) error {
    r.deviceNo = deviceNo
    r.Set("device_no", deviceNo)
    return nil
}

// DeviceNo Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetDeviceNo() string {
    return r.deviceNo
}
// AntiFakeCode Setter
// 校验码
func (r *AlibabaEinvoicePartnerUploadRequest) SetAntiFakeCode(antiFakeCode string) error {
    r.antiFakeCode = antiFakeCode
    r.Set("anti_fake_code", antiFakeCode)
    return nil
}

// AntiFakeCode Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetAntiFakeCode() string {
    return r.antiFakeCode
}
// FileDataType Setter
// 发票类型，upload_type=0且invoiceKind=0电子发票时必填，暂时只支持pdf
func (r *AlibabaEinvoicePartnerUploadRequest) SetFileDataType(fileDataType string) error {
    r.fileDataType = fileDataType
    r.Set("file_data_type", fileDataType)
    return nil
}

// FileDataType Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetFileDataType() string {
    return r.fileDataType
}
// ReqIndex Setter
// 原蓝票的reqIndex
func (r *AlibabaEinvoicePartnerUploadRequest) SetReqIndex(reqIndex string) error {
    r.reqIndex = reqIndex
    r.Set("req_index", reqIndex)
    return nil
}

// ReqIndex Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetReqIndex() string {
    return r.reqIndex
}
// InvoiceKind Setter
// 发票种类，0=电子发票，1=纸质普票，2=纸质专票
func (r *AlibabaEinvoicePartnerUploadRequest) SetInvoiceKind(invoiceKind int64) error {
    r.invoiceKind = invoiceKind
    r.Set("invoice_kind", invoiceKind)
    return nil
}

// InvoiceKind Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetInvoiceKind() int64 {
    return r.invoiceKind
}
// UploadType Setter
// 上传的类型，0=冲红上传，1=作废上传
func (r *AlibabaEinvoicePartnerUploadRequest) SetUploadType(uploadType int64) error {
    r.uploadType = uploadType
    r.Set("upload_type", uploadType)
    return nil
}

// UploadType Getter
func (r AlibabaEinvoicePartnerUploadRequest) GetUploadType() int64 {
    return r.uploadType
}
