package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票中台-发票文件下载地址查询 API请求
alibaba.einvoice.prod.result.fileurl.get

发票文件下载地址查询，外部ISV通过该接口可以查对应发票文件
*/
type AlibabaEinvoiceProdResultFileurlGetRequest struct {
    model.Params
    // 业务平台商户ID/卖家用户ID
    _platformUserId   string
    // 发票号码
    _invoiceNo   string
    // 发票代码
    _invoiceCode   string
    // 发票文件类型，小写，pdf/ofd/jpg
    _fileType   string
    // 业务平台code, 由发票中台分配
    _platformCode   string
}

// 初始化AlibabaEinvoiceProdResultFileurlGetRequest对象
func NewAlibabaEinvoiceProdResultFileurlGetRequest() *AlibabaEinvoiceProdResultFileurlGetRequest{
    return &AlibabaEinvoiceProdResultFileurlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.prod.result.fileurl.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlatformUserId Setter
// 业务平台商户ID/卖家用户ID
func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetPlatformUserId(_platformUserId string) error {
    r._platformUserId = _platformUserId
    r.Set("platform_user_id", _platformUserId)
    return nil
}

// PlatformUserId Getter
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetPlatformUserId() string {
    return r._platformUserId
}
// InvoiceNo Setter
// 发票号码
func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetInvoiceNo(_invoiceNo string) error {
    r._invoiceNo = _invoiceNo
    r.Set("invoice_no", _invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetInvoiceNo() string {
    return r._invoiceNo
}
// InvoiceCode Setter
// 发票代码
func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetInvoiceCode(_invoiceCode string) error {
    r._invoiceCode = _invoiceCode
    r.Set("invoice_code", _invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetInvoiceCode() string {
    return r._invoiceCode
}
// FileType Setter
// 发票文件类型，小写，pdf/ofd/jpg
func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetFileType(_fileType string) error {
    r._fileType = _fileType
    r.Set("file_type", _fileType)
    return nil
}

// FileType Getter
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetFileType() string {
    return r._fileType
}
// PlatformCode Setter
// 业务平台code, 由发票中台分配
func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetPlatformCode(_platformCode string) error {
    r._platformCode = _platformCode
    r.Set("platform_code", _platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetPlatformCode() string {
    return r._platformCode
}
