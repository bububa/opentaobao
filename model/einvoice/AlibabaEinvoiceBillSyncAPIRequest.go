package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicebillsyncAPIRequest 结算单同步 API请求
// alibaba.einvoice.bill.sync
//
// 电子发票业务，服务商同步结算单，包括结算单的增删改功能。最终用于开发票
type AlibabaeinvoicebillsyncAPIRequest struct {
	model.Params
	// 结算商品单明细列表
	_invoiceItems []BillItemDo
	// 结算单订单日期
	_orderDate string
	// 店铺名称，与后台店铺名称保持一致
	_shopName string
	// 税务登记证号
	_payeeRegisterNo string
	// 结算单可开票总金额（不填=sumPrice），小数点后2两位
	_invoicePrice string
	// 结算单订单ID
	_orderId string
	// 结算单总价格，小数点后2两位
	_sumPrice string
	// 调用平台，用于区分同一个税号下多个店铺来源["TB:淘宝","ALIPAY:支付宝","TM:天猫","JD:京东","DD:当当","PP:拍拍","YX:易讯","EBAY:ebay","QQ:QQ网购","AMAZON:亚马逊","SN:苏宁","GM:国美","WPH:唯品会","JM:聚美","LF:乐蜂","MGJ:蘑菇街","JS:聚尚","PX:拍鞋","YT:银泰","YHD:1号店","VANCL:凡客","YL:邮乐","YG:优购","1688:阿里巴巴","POS:POS门店","ELEME:饿了么","OTHER:其他"]
	_platform string
	// 品牌名称，不填默认=shop_name
	_brandName string
	// 开票店铺的平台，默认等于platform
	_shopPlatform string
	// 结算单同步操作：=1插入，=2更新，=3废弃删除
	_status int64
	// 生成二维码参数，若不需要生成二维码，则不填
	_qrcode *QrCodeDo
}

// NewAlibabaeinvoicebillsyncRequest 初始化AlibabaeinvoicebillsyncAPIRequest对象
func NewAlibabaeinvoicebillsyncRequest() *AlibabaeinvoicebillsyncAPIRequest {
	return &AlibabaeinvoicebillsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicebillsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.bill.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicebillsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicebillsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceItems is InvoiceItems Setter
// 结算商品单明细列表
func (r *AlibabaeinvoicebillsyncAPIRequest) SetInvoiceItems(_invoiceItems []BillItemDo) error {
	r._invoiceItems = _invoiceItems
	r.Set("invoice_items", _invoiceItems)
	return nil
}

// GetInvoiceItems InvoiceItems Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetInvoiceItems() []BillItemDo {
	return r._invoiceItems
}

// SetOrderDate is OrderDate Setter
// 结算单订单日期
func (r *AlibabaeinvoicebillsyncAPIRequest) SetOrderDate(_orderDate string) error {
	r._orderDate = _orderDate
	r.Set("order_date", _orderDate)
	return nil
}

// GetOrderDate OrderDate Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetOrderDate() string {
	return r._orderDate
}

// SetShopName is ShopName Setter
// 店铺名称，与后台店铺名称保持一致
func (r *AlibabaeinvoicebillsyncAPIRequest) SetShopName(_shopName string) error {
	r._shopName = _shopName
	r.Set("shop_name", _shopName)
	return nil
}

// GetShopName ShopName Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetShopName() string {
	return r._shopName
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 税务登记证号
func (r *AlibabaeinvoicebillsyncAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetInvoicePrice is InvoicePrice Setter
// 结算单可开票总金额（不填=sumPrice），小数点后2两位
func (r *AlibabaeinvoicebillsyncAPIRequest) SetInvoicePrice(_invoicePrice string) error {
	r._invoicePrice = _invoicePrice
	r.Set("invoice_price", _invoicePrice)
	return nil
}

// GetInvoicePrice InvoicePrice Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetInvoicePrice() string {
	return r._invoicePrice
}

// SetOrderId is OrderId Setter
// 结算单订单ID
func (r *AlibabaeinvoicebillsyncAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetSumPrice is SumPrice Setter
// 结算单总价格，小数点后2两位
func (r *AlibabaeinvoicebillsyncAPIRequest) SetSumPrice(_sumPrice string) error {
	r._sumPrice = _sumPrice
	r.Set("sum_price", _sumPrice)
	return nil
}

// GetSumPrice SumPrice Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetSumPrice() string {
	return r._sumPrice
}

// SetPlatform is Platform Setter
// 调用平台，用于区分同一个税号下多个店铺来源[&#34;TB:淘宝&#34;,&#34;ALIPAY:支付宝&#34;,&#34;TM:天猫&#34;,&#34;JD:京东&#34;,&#34;DD:当当&#34;,&#34;PP:拍拍&#34;,&#34;YX:易讯&#34;,&#34;EBAY:ebay&#34;,&#34;QQ:QQ网购&#34;,&#34;AMAZON:亚马逊&#34;,&#34;SN:苏宁&#34;,&#34;GM:国美&#34;,&#34;WPH:唯品会&#34;,&#34;JM:聚美&#34;,&#34;LF:乐蜂&#34;,&#34;MGJ:蘑菇街&#34;,&#34;JS:聚尚&#34;,&#34;PX:拍鞋&#34;,&#34;YT:银泰&#34;,&#34;YHD:1号店&#34;,&#34;VANCL:凡客&#34;,&#34;YL:邮乐&#34;,&#34;YG:优购&#34;,&#34;1688:阿里巴巴&#34;,&#34;POS:POS门店&#34;,&#34;ELEME:饿了么&#34;,&#34;OTHER:其他&#34;]
func (r *AlibabaeinvoicebillsyncAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetPlatform() string {
	return r._platform
}

// SetBrandName is BrandName Setter
// 品牌名称，不填默认=shop_name
func (r *AlibabaeinvoicebillsyncAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// GetBrandName BrandName Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetBrandName() string {
	return r._brandName
}

// SetShopPlatform is ShopPlatform Setter
// 开票店铺的平台，默认等于platform
func (r *AlibabaeinvoicebillsyncAPIRequest) SetShopPlatform(_shopPlatform string) error {
	r._shopPlatform = _shopPlatform
	r.Set("shop_platform", _shopPlatform)
	return nil
}

// GetShopPlatform ShopPlatform Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetShopPlatform() string {
	return r._shopPlatform
}

// SetStatus is Status Setter
// 结算单同步操作：=1插入，=2更新，=3废弃删除
func (r *AlibabaeinvoicebillsyncAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetStatus() int64 {
	return r._status
}

// SetQrcode is Qrcode Setter
// 生成二维码参数，若不需要生成二维码，则不填
func (r *AlibabaeinvoicebillsyncAPIRequest) SetQrcode(_qrcode *QrCodeDo) error {
	r._qrcode = _qrcode
	r.Set("qrcode", _qrcode)
	return nil
}

// GetQrcode Qrcode Getter
func (r AlibabaeinvoicebillsyncAPIRequest) GetQrcode() *QrCodeDo {
	return r._qrcode
}
