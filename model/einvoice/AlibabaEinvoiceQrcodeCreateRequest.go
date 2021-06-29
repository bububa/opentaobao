package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码开票二维码生成 API请求
alibaba.einvoice.qrcode.create

扫码开票功能中的二维码生成接口，pos机等发起请求生成二维码
*/
type AlibabaEinvoiceQrcodeCreateRequest struct {
    model.Params
    // 发票商品明细
    _invoiceItems   []BillItemDo
    // 收款方税务登记证号
    _payeeRegisterNo   string
    // 开票的订单号，同结算单订单号
    _orderId   string
    // 开票金额
    _sumPrice   string
    // 请求方唯一标识ID，例如POS机编码
    _sourceId   string
    // 二维码图片中间的logo
    _qrLogo   string
    // 二维码图片宽度，默认=450
    _width   int64
    // 二维码图片高度，默认=450
    _height   int64
    // 二维码返回类型：1=二维码背后的URL，2=二维码图片CDN URL，3=二维码二进制数据流
    _qrType   int64
    // 请求方ERP系统平台，同结算单同步的platform，不填默认=platform
    _platform   string
}

// 初始化AlibabaEinvoiceQrcodeCreateRequest对象
func NewAlibabaEinvoiceQrcodeCreateRequest() *AlibabaEinvoiceQrcodeCreateRequest{
    return &AlibabaEinvoiceQrcodeCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceQrcodeCreateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.qrcode.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceQrcodeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceItems Setter
// 发票商品明细
func (r *AlibabaEinvoiceQrcodeCreateRequest) SetInvoiceItems(_invoiceItems []BillItemDo) error {
    r._invoiceItems = _invoiceItems
    r.Set("invoice_items", _invoiceItems)
    return nil
}

// InvoiceItems Getter
func (r AlibabaEinvoiceQrcodeCreateRequest) GetInvoiceItems() []BillItemDo {
    return r._invoiceItems
}
// PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceQrcodeCreateRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceQrcodeCreateRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// OrderId Setter
// 开票的订单号，同结算单订单号
func (r *AlibabaEinvoiceQrcodeCreateRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEinvoiceQrcodeCreateRequest) GetOrderId() string {
    return r._orderId
}
// SumPrice Setter
// 开票金额
func (r *AlibabaEinvoiceQrcodeCreateRequest) SetSumPrice(_sumPrice string) error {
    r._sumPrice = _sumPrice
    r.Set("sum_price", _sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceQrcodeCreateRequest) GetSumPrice() string {
    return r._sumPrice
}
// SourceId Setter
// 请求方唯一标识ID，例如POS机编码
func (r *AlibabaEinvoiceQrcodeCreateRequest) SetSourceId(_sourceId string) error {
    r._sourceId = _sourceId
    r.Set("source_id", _sourceId)
    return nil
}

// SourceId Getter
func (r AlibabaEinvoiceQrcodeCreateRequest) GetSourceId() string {
    return r._sourceId
}
// QrLogo Setter
// 二维码图片中间的logo
func (r *AlibabaEinvoiceQrcodeCreateRequest) SetQrLogo(_qrLogo string) error {
    r._qrLogo = _qrLogo
    r.Set("qr_logo", _qrLogo)
    return nil
}

// QrLogo Getter
func (r AlibabaEinvoiceQrcodeCreateRequest) GetQrLogo() string {
    return r._qrLogo
}
// Width Setter
// 二维码图片宽度，默认=450
func (r *AlibabaEinvoiceQrcodeCreateRequest) SetWidth(_width int64) error {
    r._width = _width
    r.Set("width", _width)
    return nil
}

// Width Getter
func (r AlibabaEinvoiceQrcodeCreateRequest) GetWidth() int64 {
    return r._width
}
// Height Setter
// 二维码图片高度，默认=450
func (r *AlibabaEinvoiceQrcodeCreateRequest) SetHeight(_height int64) error {
    r._height = _height
    r.Set("height", _height)
    return nil
}

// Height Getter
func (r AlibabaEinvoiceQrcodeCreateRequest) GetHeight() int64 {
    return r._height
}
// QrType Setter
// 二维码返回类型：1=二维码背后的URL，2=二维码图片CDN URL，3=二维码二进制数据流
func (r *AlibabaEinvoiceQrcodeCreateRequest) SetQrType(_qrType int64) error {
    r._qrType = _qrType
    r.Set("qr_type", _qrType)
    return nil
}

// QrType Getter
func (r AlibabaEinvoiceQrcodeCreateRequest) GetQrType() int64 {
    return r._qrType
}
// Platform Setter
// 请求方ERP系统平台，同结算单同步的platform，不填默认=platform
func (r *AlibabaEinvoiceQrcodeCreateRequest) SetPlatform(_platform string) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r AlibabaEinvoiceQrcodeCreateRequest) GetPlatform() string {
    return r._platform
}
