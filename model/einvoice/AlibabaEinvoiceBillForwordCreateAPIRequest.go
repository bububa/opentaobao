package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码开票结算单同步前开发票 API请求
alibaba.einvoice.bill.forword.create

扫码开票结算单同步前开发票，会将数据同步到结算单中
*/
type AlibabaEinvoiceBillForwordCreateAPIRequest struct {
    model.Params
    // 发票商品明细
    _invoiceItems   []BillItemDo
    // 结算单订单ID
    _orderId   string
    // 收款方税号
    _payeeRegisterNo   string
    // 调用平台，用以区分不同的订单ID，不填默认为default
    _platform   string
    // 开票流水号，若不填则系统默认生成
    _seriNo   string
    // 店铺名称
    _shopName   string
    // 开票金额，和明细累计总金额需相同
    _sumPrice   string
    // 付款方地址
    _payerAddress   string
    // 付款方银行账户
    _payerBankaccount   string
    // 付款方开票邮件通知邮箱
    _payerEmail   string
    // 付款方发票抬头
    _payerName   string
    // 企业电话
    _payerPhone   string
    // 付款方税号
    _payerRegisterNo   string
    // 个人电话，接收发票通知
    _phoneNumber   string
    // 企业或个人抬头发票，0=个人，1=企业。默认=1
    _businessType   int64
}

// 初始化AlibabaEinvoiceBillForwordCreateAPIRequest对象
func NewAlibabaEinvoiceBillForwordCreateRequest() *AlibabaEinvoiceBillForwordCreateAPIRequest{
    return &AlibabaEinvoiceBillForwordCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.einvoice.bill.forword.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceItems Setter
// 发票商品明细
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetInvoiceItems(_invoiceItems []BillItemDo) error {
    r._invoiceItems = _invoiceItems
    r.Set("invoice_items", _invoiceItems)
    return nil
}

// InvoiceItems Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetInvoiceItems() []BillItemDo {
    return r._invoiceItems
}
// OrderId Setter
// 结算单订单ID
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetOrderId() string {
    return r._orderId
}
// PayeeRegisterNo Setter
// 收款方税号
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// Platform Setter
// 调用平台，用以区分不同的订单ID，不填默认为default
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetPlatform(_platform string) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetPlatform() string {
    return r._platform
}
// SeriNo Setter
// 开票流水号，若不填则系统默认生成
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetSeriNo(_seriNo string) error {
    r._seriNo = _seriNo
    r.Set("seri_no", _seriNo)
    return nil
}

// SeriNo Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetSeriNo() string {
    return r._seriNo
}
// ShopName Setter
// 店铺名称
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetShopName(_shopName string) error {
    r._shopName = _shopName
    r.Set("shop_name", _shopName)
    return nil
}

// ShopName Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetShopName() string {
    return r._shopName
}
// SumPrice Setter
// 开票金额，和明细累计总金额需相同
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetSumPrice(_sumPrice string) error {
    r._sumPrice = _sumPrice
    r.Set("sum_price", _sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetSumPrice() string {
    return r._sumPrice
}
// PayerAddress Setter
// 付款方地址
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetPayerAddress(_payerAddress string) error {
    r._payerAddress = _payerAddress
    r.Set("payer_address", _payerAddress)
    return nil
}

// PayerAddress Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetPayerAddress() string {
    return r._payerAddress
}
// PayerBankaccount Setter
// 付款方银行账户
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetPayerBankaccount(_payerBankaccount string) error {
    r._payerBankaccount = _payerBankaccount
    r.Set("payer_bankaccount", _payerBankaccount)
    return nil
}

// PayerBankaccount Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetPayerBankaccount() string {
    return r._payerBankaccount
}
// PayerEmail Setter
// 付款方开票邮件通知邮箱
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetPayerEmail(_payerEmail string) error {
    r._payerEmail = _payerEmail
    r.Set("payer_email", _payerEmail)
    return nil
}

// PayerEmail Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetPayerEmail() string {
    return r._payerEmail
}
// PayerName Setter
// 付款方发票抬头
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetPayerName(_payerName string) error {
    r._payerName = _payerName
    r.Set("payer_name", _payerName)
    return nil
}

// PayerName Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetPayerName() string {
    return r._payerName
}
// PayerPhone Setter
// 企业电话
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetPayerPhone(_payerPhone string) error {
    r._payerPhone = _payerPhone
    r.Set("payer_phone", _payerPhone)
    return nil
}

// PayerPhone Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetPayerPhone() string {
    return r._payerPhone
}
// PayerRegisterNo Setter
// 付款方税号
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetPayerRegisterNo(_payerRegisterNo string) error {
    r._payerRegisterNo = _payerRegisterNo
    r.Set("payer_register_no", _payerRegisterNo)
    return nil
}

// PayerRegisterNo Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetPayerRegisterNo() string {
    return r._payerRegisterNo
}
// PhoneNumber Setter
// 个人电话，接收发票通知
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetPhoneNumber(_phoneNumber string) error {
    r._phoneNumber = _phoneNumber
    r.Set("phone_number", _phoneNumber)
    return nil
}

// PhoneNumber Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetPhoneNumber() string {
    return r._phoneNumber
}
// BusinessType Setter
// 企业或个人抬头发票，0=个人，1=企业。默认=1
func (r *AlibabaEinvoiceBillForwordCreateAPIRequest) SetBusinessType(_businessType int64) error {
    r._businessType = _businessType
    r.Set("business_type", _businessType)
    return nil
}

// BusinessType Getter
func (r AlibabaEinvoiceBillForwordCreateAPIRequest) GetBusinessType() int64 {
    return r._businessType
}
