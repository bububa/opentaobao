package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传发票ocr的结果 APIRequest
alibaba.einvoice.income.ocr.return

服务商回传发票ocr的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的ocr回传
*/
type AlibabaEinvoiceIncomeOcrReturnRequest struct {
    model.Params

    // 校验码，ocr结果为普票，success=true并且invoiceKind=1时必填
    checksum   string 

    // 错误码，success=false是必填
    errorCode   string 

    // 错误消息，success=false是必填
    errorMessage   string 

    // 发票ocr影像文件，type=1时必填
    imageData   []*model.File 

    // 发票ocr影像编号，type=1时必填
    imageId   string 

    // 发票代码，success=true时必填
    invoiceCode   string 

    // 开票日期，格式为yyyy-MM-dd，success=true时必填
    invoiceDate   string 

    // 发票种类，1=普票，2=专票，success=true时必填
    invoiceKind   int64 

    // 发票号码，success=true时必填
    invoiceNo   string 

    // 开票请求标识，扫描驱动回传type=1时填批次号
    reqIndex   string 

    // ocr结果，true=成功，false=失败
    success   bool 

    // 不含税金额，ocr结果为专票，success=true并且invoiceKind=2时必填
    sumPrice   string 

    // 请求类型，0=阿里主动发起的cor，1=扫描驱动服务商主动回传ocr结果
    type   int64 

}

func NewAlibabaEinvoiceIncomeOcrReturnRequest() *AlibabaEinvoiceIncomeOcrReturnRequest{
    return &AlibabaEinvoiceIncomeOcrReturnRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.income.ocr.return"
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetChecksum(checksum string) error {
    r.checksum = checksum
    r.Set("checksum", checksum)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetChecksum() string {
    return r.checksum
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetErrorCode(errorCode string) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetErrorCode() string {
    return r.errorCode
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetErrorMessage(errorMessage string) error {
    r.errorMessage = errorMessage
    r.Set("error_message", errorMessage)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetErrorMessage() string {
    return r.errorMessage
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetImageData(imageData []*model.File) error {
    r.imageData = imageData
    r.Set("image_data", imageData)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetImageData() []*model.File {
    return r.imageData
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetImageId(imageId string) error {
    r.imageId = imageId
    r.Set("image_id", imageId)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetImageId() string {
    return r.imageId
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceCode(invoiceCode string) error {
    r.invoiceCode = invoiceCode
    r.Set("invoice_code", invoiceCode)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceCode() string {
    return r.invoiceCode
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceDate(invoiceDate string) error {
    r.invoiceDate = invoiceDate
    r.Set("invoice_date", invoiceDate)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceDate() string {
    return r.invoiceDate
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceKind(invoiceKind int64) error {
    r.invoiceKind = invoiceKind
    r.Set("invoice_kind", invoiceKind)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceKind() int64 {
    return r.invoiceKind
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetInvoiceNo(invoiceNo string) error {
    r.invoiceNo = invoiceNo
    r.Set("invoice_no", invoiceNo)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetInvoiceNo() string {
    return r.invoiceNo
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetReqIndex(reqIndex string) error {
    r.reqIndex = reqIndex
    r.Set("req_index", reqIndex)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetReqIndex() string {
    return r.reqIndex
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetSuccess() bool {
    return r.success
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetSumPrice(sumPrice string) error {
    r.sumPrice = sumPrice
    r.Set("sum_price", sumPrice)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetSumPrice() string {
    return r.sumPrice
}

func (r *AlibabaEinvoiceIncomeOcrReturnRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaEinvoiceIncomeOcrReturnRequest) GetType() int64 {
    return r.type
}

