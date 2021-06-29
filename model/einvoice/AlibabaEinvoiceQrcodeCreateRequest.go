package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码开票二维码生成 APIRequest
alibaba.einvoice.qrcode.create

扫码开票功能中的二维码生成接口，pos机等发起请求生成二维码
*/
type AlibabaEinvoiceQrcodeCreateRequest struct {
    model.Params

    // 发票商品明细
    invoiceItems   []BillItemDo 

    // 收款方税务登记证号
    payeeRegisterNo   string 

    // 开票的订单号，同结算单订单号
    orderId   string 

    // 开票金额
    sumPrice   string 

    // 请求方唯一标识ID，例如POS机编码
    sourceId   string 

    // 二维码图片中间的logo
    qrLogo   string 

    // 二维码图片宽度，默认=450
    width   int64 

    // 二维码图片高度，默认=450
    height   int64 

    // 二维码返回类型：1=二维码背后的URL，2=二维码图片CDN URL，3=二维码二进制数据流
    qrType   int64 

    // 请求方ERP系统平台，同结算单同步的platform，不填默认=platform
    platform   string 

}

func NewAlibabaEinvoiceQrcodeCreateRequest() *AlibabaEinvoiceQrcodeCreateRequest{
    return &AlibabaEinvoiceQrcodeCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.qrcode.create"
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceQrcodeCreateRequest) SetInvoiceItems(invoiceItems []BillItemDo) error {
    r.invoiceItems = invoiceItems
    r.Set("invoice_items", invoiceItems)
    return nil
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetInvoiceItems() []BillItemDo {
    return r.invoiceItems
}

func (r *AlibabaEinvoiceQrcodeCreateRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoiceQrcodeCreateRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaEinvoiceQrcodeCreateRequest) SetSumPrice(sumPrice string) error {
    r.sumPrice = sumPrice
    r.Set("sum_price", sumPrice)
    return nil
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetSumPrice() string {
    return r.sumPrice
}

func (r *AlibabaEinvoiceQrcodeCreateRequest) SetSourceId(sourceId string) error {
    r.sourceId = sourceId
    r.Set("source_id", sourceId)
    return nil
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetSourceId() string {
    return r.sourceId
}

func (r *AlibabaEinvoiceQrcodeCreateRequest) SetQrLogo(qrLogo string) error {
    r.qrLogo = qrLogo
    r.Set("qr_logo", qrLogo)
    return nil
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetQrLogo() string {
    return r.qrLogo
}

func (r *AlibabaEinvoiceQrcodeCreateRequest) SetWidth(width int64) error {
    r.width = width
    r.Set("width", width)
    return nil
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetWidth() int64 {
    return r.width
}

func (r *AlibabaEinvoiceQrcodeCreateRequest) SetHeight(height int64) error {
    r.height = height
    r.Set("height", height)
    return nil
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetHeight() int64 {
    return r.height
}

func (r *AlibabaEinvoiceQrcodeCreateRequest) SetQrType(qrType int64) error {
    r.qrType = qrType
    r.Set("qr_type", qrType)
    return nil
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetQrType() int64 {
    return r.qrType
}

func (r *AlibabaEinvoiceQrcodeCreateRequest) SetPlatform(platform string) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r AlibabaEinvoiceQrcodeCreateRequest) GetPlatform() string {
    return r.platform
}

