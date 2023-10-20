package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceQrcodeCreateAPIRequest 扫码开票二维码生成 API请求
// alibaba.einvoice.qrcode.create
//
// 扫码开票功能中的二维码生成接口，pos机等发起请求生成二维码
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
	// 请求方ERP系统平台，同结算单同步的platform，不填默认=platform
	_platform string
	// 二维码图片中间的logo
	_qrLogo string
	// 二维码图片宽度，默认=450
	_width int64
	// 二维码图片高度，默认=450
	_height int64
	// 二维码返回类型：1=二维码背后的URL，2=二维码图片CDN URL，3=二维码二进制数据流
	_qrType int64
}

// NewAlibabaEinvoiceQrcodeCreateRequest 初始化AlibabaEinvoiceQrcodeCreateAPIRequest对象
func NewAlibabaEinvoiceQrcodeCreateRequest() *AlibabaEinvoiceQrcodeCreateAPIRequest {
	return &AlibabaEinvoiceQrcodeCreateAPIRequest{
		Params: model.NewParams(10),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) Reset() {
	r._invoiceItems = r._invoiceItems[:0]
	r._payeeRegisterNo = ""
	r._orderId = ""
	r._sumPrice = ""
	r._sourceId = ""
	r._platform = ""
	r._qrLogo = ""
	r._width = 0
	r._height = 0
	r._qrType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.qrcode.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceItems is InvoiceItems Setter
// 发票商品明细
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetInvoiceItems(_invoiceItems []BillItemDo) error {
	r._invoiceItems = _invoiceItems
	r.Set("invoice_items", _invoiceItems)
	return nil
}

// GetInvoiceItems InvoiceItems Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetInvoiceItems() []BillItemDo {
	return r._invoiceItems
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetOrderId is OrderId Setter
// 开票的订单号，同结算单订单号
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetSumPrice is SumPrice Setter
// 开票金额
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetSumPrice(_sumPrice string) error {
	r._sumPrice = _sumPrice
	r.Set("sum_price", _sumPrice)
	return nil
}

// GetSumPrice SumPrice Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetSumPrice() string {
	return r._sumPrice
}

// SetSourceId is SourceId Setter
// 请求方唯一标识ID，例如POS机编码
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetSourceId(_sourceId string) error {
	r._sourceId = _sourceId
	r.Set("source_id", _sourceId)
	return nil
}

// GetSourceId SourceId Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetSourceId() string {
	return r._sourceId
}

// SetPlatform is Platform Setter
// 请求方ERP系统平台，同结算单同步的platform，不填默认=platform
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetPlatform() string {
	return r._platform
}

// SetQrLogo is QrLogo Setter
// 二维码图片中间的logo
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetQrLogo(_qrLogo string) error {
	r._qrLogo = _qrLogo
	r.Set("qr_logo", _qrLogo)
	return nil
}

// GetQrLogo QrLogo Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetQrLogo() string {
	return r._qrLogo
}

// SetWidth is Width Setter
// 二维码图片宽度，默认=450
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetWidth(_width int64) error {
	r._width = _width
	r.Set("width", _width)
	return nil
}

// GetWidth Width Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetWidth() int64 {
	return r._width
}

// SetHeight is Height Setter
// 二维码图片高度，默认=450
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// GetHeight Height Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetHeight() int64 {
	return r._height
}

// SetQrType is QrType Setter
// 二维码返回类型：1=二维码背后的URL，2=二维码图片CDN URL，3=二维码二进制数据流
func (r *AlibabaEinvoiceQrcodeCreateAPIRequest) SetQrType(_qrType int64) error {
	r._qrType = _qrType
	r.Set("qr_type", _qrType)
	return nil
}

// GetQrType QrType Getter
func (r AlibabaEinvoiceQrcodeCreateAPIRequest) GetQrType() int64 {
	return r._qrType
}

var poolAlibabaEinvoiceQrcodeCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceQrcodeCreateRequest()
	},
}

// GetAlibabaEinvoiceQrcodeCreateRequest 从 sync.Pool 获取 AlibabaEinvoiceQrcodeCreateAPIRequest
func GetAlibabaEinvoiceQrcodeCreateAPIRequest() *AlibabaEinvoiceQrcodeCreateAPIRequest {
	return poolAlibabaEinvoiceQrcodeCreateAPIRequest.Get().(*AlibabaEinvoiceQrcodeCreateAPIRequest)
}

// ReleaseAlibabaEinvoiceQrcodeCreateAPIRequest 将 AlibabaEinvoiceQrcodeCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceQrcodeCreateAPIRequest(v *AlibabaEinvoiceQrcodeCreateAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceQrcodeCreateAPIRequest.Put(v)
}
