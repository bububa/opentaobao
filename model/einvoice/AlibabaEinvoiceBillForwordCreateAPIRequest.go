package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicebillforwordcreateAPIRequest 扫码开票结算单同步前开发票 API请求
// alibaba.einvoice.bill.forword.create
//
// 扫码开票结算单同步前开发票，会将数据同步到结算单中
type AlibabaeinvoicebillforwordcreateAPIRequest struct {
	model.Params
	// 发票商品明细
	_invoiceItems []BillItemDo
	// 结算单订单ID
	_orderId string
	// 收款方税号
	_payeeRegisterNo string
	// 调用平台，用以区分不同的订单ID，不填默认为default
	_platform string
	// 开票流水号，若不填则系统默认生成
	_seriNo string
	// 店铺名称
	_shopName string
	// 开票金额，和明细累计总金额需相同
	_sumPrice string
	// 付款方地址
	_payerAddress string
	// 付款方银行账户
	_payerBankaccount string
	// 付款方开票邮件通知邮箱
	_payerEmail string
	// 付款方发票抬头
	_payerName string
	// 企业电话
	_payerPhone string
	// 付款方税号
	_payerRegisterNo string
	// 个人电话，接收发票通知
	_phoneNumber string
	// 企业或个人抬头发票，0=个人，1=企业。默认=1
	_businessType int64
}

// NewAlibabaeinvoicebillforwordcreateRequest 初始化AlibabaeinvoicebillforwordcreateAPIRequest对象
func NewAlibabaeinvoicebillforwordcreateRequest() *AlibabaeinvoicebillforwordcreateAPIRequest {
	return &AlibabaeinvoicebillforwordcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.bill.forword.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceItems is InvoiceItems Setter
// 发票商品明细
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetInvoiceItems(_invoiceItems []BillItemDo) error {
	r._invoiceItems = _invoiceItems
	r.Set("invoice_items", _invoiceItems)
	return nil
}

// GetInvoiceItems InvoiceItems Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetInvoiceItems() []BillItemDo {
	return r._invoiceItems
}

// SetOrderId is OrderId Setter
// 结算单订单ID
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 收款方税号
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetPlatform is Platform Setter
// 调用平台，用以区分不同的订单ID，不填默认为default
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetPlatform() string {
	return r._platform
}

// SetSeriNo is SeriNo Setter
// 开票流水号，若不填则系统默认生成
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetSeriNo(_seriNo string) error {
	r._seriNo = _seriNo
	r.Set("seri_no", _seriNo)
	return nil
}

// GetSeriNo SeriNo Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetSeriNo() string {
	return r._seriNo
}

// SetShopName is ShopName Setter
// 店铺名称
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetShopName(_shopName string) error {
	r._shopName = _shopName
	r.Set("shop_name", _shopName)
	return nil
}

// GetShopName ShopName Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetShopName() string {
	return r._shopName
}

// SetSumPrice is SumPrice Setter
// 开票金额，和明细累计总金额需相同
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetSumPrice(_sumPrice string) error {
	r._sumPrice = _sumPrice
	r.Set("sum_price", _sumPrice)
	return nil
}

// GetSumPrice SumPrice Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetSumPrice() string {
	return r._sumPrice
}

// SetPayerAddress is PayerAddress Setter
// 付款方地址
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetPayerAddress(_payerAddress string) error {
	r._payerAddress = _payerAddress
	r.Set("payer_address", _payerAddress)
	return nil
}

// GetPayerAddress PayerAddress Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetPayerAddress() string {
	return r._payerAddress
}

// SetPayerBankaccount is PayerBankaccount Setter
// 付款方银行账户
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetPayerBankaccount(_payerBankaccount string) error {
	r._payerBankaccount = _payerBankaccount
	r.Set("payer_bankaccount", _payerBankaccount)
	return nil
}

// GetPayerBankaccount PayerBankaccount Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetPayerBankaccount() string {
	return r._payerBankaccount
}

// SetPayerEmail is PayerEmail Setter
// 付款方开票邮件通知邮箱
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetPayerEmail(_payerEmail string) error {
	r._payerEmail = _payerEmail
	r.Set("payer_email", _payerEmail)
	return nil
}

// GetPayerEmail PayerEmail Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetPayerEmail() string {
	return r._payerEmail
}

// SetPayerName is PayerName Setter
// 付款方发票抬头
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetPayerName(_payerName string) error {
	r._payerName = _payerName
	r.Set("payer_name", _payerName)
	return nil
}

// GetPayerName PayerName Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetPayerName() string {
	return r._payerName
}

// SetPayerPhone is PayerPhone Setter
// 企业电话
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetPayerPhone(_payerPhone string) error {
	r._payerPhone = _payerPhone
	r.Set("payer_phone", _payerPhone)
	return nil
}

// GetPayerPhone PayerPhone Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetPayerPhone() string {
	return r._payerPhone
}

// SetPayerRegisterNo is PayerRegisterNo Setter
// 付款方税号
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetPayerRegisterNo(_payerRegisterNo string) error {
	r._payerRegisterNo = _payerRegisterNo
	r.Set("payer_register_no", _payerRegisterNo)
	return nil
}

// GetPayerRegisterNo PayerRegisterNo Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetPayerRegisterNo() string {
	return r._payerRegisterNo
}

// SetPhoneNumber is PhoneNumber Setter
// 个人电话，接收发票通知
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetPhoneNumber(_phoneNumber string) error {
	r._phoneNumber = _phoneNumber
	r.Set("phone_number", _phoneNumber)
	return nil
}

// GetPhoneNumber PhoneNumber Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetPhoneNumber() string {
	return r._phoneNumber
}

// SetBusinessType is BusinessType Setter
// 企业或个人抬头发票，0=个人，1=企业。默认=1
func (r *AlibabaeinvoicebillforwordcreateAPIRequest) SetBusinessType(_businessType int64) error {
	r._businessType = _businessType
	r.Set("business_type", _businessType)
	return nil
}

// GetBusinessType BusinessType Getter
func (r AlibabaeinvoicebillforwordcreateAPIRequest) GetBusinessType() int64 {
	return r._businessType
}
