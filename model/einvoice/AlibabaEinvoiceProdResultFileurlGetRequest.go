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
    platformUserId   string
    // 发票号码
    invoiceNo   string
    // 发票代码
    invoiceCode   string
    // 发票文件类型，小写，pdf/ofd/jpg
    fileType   string
    // 业务平台code, 由发票中台分配
    platformCode   string
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
func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetPlatformUserId(platformUserId string) error {
    r.platformUserId = platformUserId
    r.Set("platform_user_id", platformUserId)
    return nil
}

// PlatformUserId Getter
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetPlatformUserId() string {
    return r.platformUserId
}
// InvoiceNo Setter
// 发票号码
func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

// InvoiceNo Getter
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetInvoiceNo() string {
    return r.invoiceNo
}
// InvoiceCode Setter
// 发票代码
func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

// InvoiceCode Getter
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetInvoiceCode() string {
    return r.invoiceCode
}
// FileType Setter
// 发票文件类型，小写，pdf/ofd/jpg
func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetFileType(fileType string) error {
    r.fileType = fileType
    r.Set("file_type", fileType)
    return nil
}

// FileType Getter
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetFileType() string {
    return r.fileType
}
// PlatformCode Setter
// 业务平台code, 由发票中台分配
func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetPlatformCode() string {
    return r.platformCode
}
