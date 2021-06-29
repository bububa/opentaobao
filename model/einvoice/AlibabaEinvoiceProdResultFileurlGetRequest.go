package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票中台-发票文件下载地址查询 APIRequest
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

func NewAlibabaEinvoiceProdResultFileurlGetRequest() *AlibabaEinvoiceProdResultFileurlGetRequest{
    return &AlibabaEinvoiceProdResultFileurlGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.prod.result.fileurl.get"
}

func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetPlatformUserId(platformUserId string) error {
    r.platformUserId = platformUserId
    r.Set("platform_user_id", platformUserId)
    return nil
}

func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetPlatformUserId() string {
    return r.platformUserId
}

func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetInvoiceNo() string {
    return r.invoiceNo
}

func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetInvoiceCode() string {
    return r.invoiceCode
}

func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetFileType(fileType string) error {
    r.fileType = fileType
    r.Set("file_type", fileType)
    return nil
}

func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetFileType() string {
    return r.fileType
}

func (r *AlibabaEinvoiceProdResultFileurlGetRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

func (r AlibabaEinvoiceProdResultFileurlGetRequest) GetPlatformCode() string {
    return r.platformCode
}

