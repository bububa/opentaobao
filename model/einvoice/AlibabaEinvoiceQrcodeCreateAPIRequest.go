package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceQrcodeCreateAPIRequest
扫码开票二维码生成 API请求
alibaba.einvoice.qrcode.create

扫码开票功能中的二维码生成接口，pos机等发起请求生成二维码 */
type AlibabaEinvoiceQrcodeCreateAPIRequest struct {
	model.Params
	// 发票商品明细
	_invoiceItems []BillItemDo
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 开票的订单号，同结算单订单号
	_orderId string
	// 开票金额
	_sumPrice string
	// 请求方唯一标识ID，例如POS机编码
	_sourceId string
	// 二维码图片中间的logo
	_qrLogo string
	// 二维码图片宽度，默认=450
	_width int64
	// 二维码图片高度，默认=450
	_height int64
	// 二维码返回类型：1=二维码背后的URL，2=二维码图片CDN URL，3=二维码二进制数据流
	_qrType int64
	// 请求方ERP系统平台，同结算单同步的platform，不填默认=platform
	_platform string
}

// NewAlibabaEinvoiceQrcodeCreateRequest 初始化AlibabaEinvoiceQrcodeCreateAPIRequest对象
func NewAlibabaEinvoiceQrcodeCreateRequest() *AlibabaEinvoiceQrcodeCreateAPIRequest {
	return &AlibabaEinvoiceQrcodeCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.qrcode.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is InvoiceItems Setter
// 发票商品明细
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetInvoiceItems(_invoiceItems []BillItemDo) error {
	r._invoiceItems = _invoiceItems
	r.Set("invoice_items", _invoiceItems)
	return nil
}

// Get InvoiceItems Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetInvoiceItems() []BillItemDo {
	return r._invoiceItems
}

// Set is PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// Get PayeeRegisterNo Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// Set is OrderId Setter
// 开票的订单号，同结算单订单号
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetOrderId() string {
	return r._orderId
}

// Set is SumPrice Setter
// 开票金额
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetSumPrice(_sumPrice string) error {
	r._sumPrice = _sumPrice
	r.Set("sum_price", _sumPrice)
	return nil
}

// Get SumPrice Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetSumPrice() string {
	return r._sumPrice
}

// Set is SourceId Setter
// 请求方唯一标识ID，例如POS机编码
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetSourceId(_sourceId string) error {
	r._sourceId = _sourceId
	r.Set("source_id", _sourceId)
	return nil
}

// Get SourceId Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetSourceId() string {
	return r._sourceId
}

// Set is QrLogo Setter
// 二维码图片中间的logo
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetQrLogo(_qrLogo string) error {
	r._qrLogo = _qrLogo
	r.Set("qr_logo", _qrLogo)
	return nil
}

// Get QrLogo Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetQrLogo() string {
	return r._qrLogo
}

// Set is Width Setter
// 二维码图片宽度，默认=450
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetWidth(_width int64) error {
	r._width = _width
	r.Set("width", _width)
	return nil
}

// Get Width Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetWidth() int64 {
	return r._width
}

// Set is Height Setter
// 二维码图片高度，默认=450
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// Get Height Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetHeight() int64 {
	return r._height
}

// Set is QrType Setter
// 二维码返回类型：1=二维码背后的URL，2=二维码图片CDN URL，3=二维码二进制数据流
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetQrType(_qrType int64) error {
	r._qrType = _qrType
	r.Set("qr_type", _qrType)
	return nil
}

// Get QrType Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetQrType() int64 {
	return r._qrType
}

// Set is Platform Setter
// 请求方ERP系统平台，同结算单同步的platform，不填默认=platform
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// Get Platform Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetPlatform() string {
	return r._platform
}
