package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
结算单同步 API请求
alibaba.einvoice.bill.sync

电子发票业务，服务商同步结算单，包括结算单的增删改功能。最终用于开发票
*/
type AlibabaEinvoiceBillSyncRequest struct {
    model.Params
    // 结算商品单明细列表
    _invoiceItems   []BillItemDo
    // 结算单同步操作：=1插入，=2更新，=3废弃删除
    _status   int64
    // 结算单订单日期
    _orderDate   string
    // 店铺名称，与后台店铺名称保持一致
    _shopName   string
    // 税务登记证号
    _payeeRegisterNo   string
    // 结算单订单ID
    _orderId   string
    // 结算单总价格，小数点后2两位
    _sumPrice   string
    // 调用平台，用于区分同一个税号下多个店铺来源["TB:淘宝","ALIPAY:支付宝","TM:天猫","JD:京东","DD:当当","PP:拍拍","YX:易讯","EBAY:ebay","QQ:QQ网购","AMAZON:亚马逊","SN:苏宁","GM:国美","WPH:唯品会","JM:聚美","LF:乐蜂","MGJ:蘑菇街","JS:聚尚","PX:拍鞋","YT:银泰","YHD:1号店","VANCL:凡客","YL:邮乐","YG:优购","1688:阿里巴巴","POS:POS门店","ELEME:饿了么","OTHER:其他"]
    _platform   string
    // 生成二维码参数，若不需要生成二维码，则不填
    _qrcode   *QrCodeDo
    // 品牌名称，不填默认=shop_name
    _brandName   string
    // 结算单可开票总金额（不填=sumPrice），小数点后2两位
    _invoicePrice   string
    // 开票店铺的平台，默认等于platform
    _shopPlatform   string
}

// 初始化AlibabaEinvoiceBillSyncRequest对象
func NewAlibabaEinvoiceBillSyncRequest() *AlibabaEinvoiceBillSyncRequest{
    return &AlibabaEinvoiceBillSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceBillSyncRequest) GetApiMethodName() string {
    return "alibaba.einvoice.bill.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceBillSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceItems Setter
// 结算商品单明细列表
func (r *AlibabaEinvoiceBillSyncRequest) SetInvoiceItems(_invoiceItems []BillItemDo) error {
    r._invoiceItems = _invoiceItems
    r.Set("invoice_items", _invoiceItems)
    return nil
}

// InvoiceItems Getter
func (r AlibabaEinvoiceBillSyncRequest) GetInvoiceItems() []BillItemDo {
    return r._invoiceItems
}
// Status Setter
// 结算单同步操作：=1插入，=2更新，=3废弃删除
func (r *AlibabaEinvoiceBillSyncRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaEinvoiceBillSyncRequest) GetStatus() int64 {
    return r._status
}
// OrderDate Setter
// 结算单订单日期
func (r *AlibabaEinvoiceBillSyncRequest) SetOrderDate(_orderDate string) error {
    r._orderDate = _orderDate
    r.Set("order_date", _orderDate)
    return nil
}

// OrderDate Getter
func (r AlibabaEinvoiceBillSyncRequest) GetOrderDate() string {
    return r._orderDate
}
// ShopName Setter
// 店铺名称，与后台店铺名称保持一致
func (r *AlibabaEinvoiceBillSyncRequest) SetShopName(_shopName string) error {
    r._shopName = _shopName
    r.Set("shop_name", _shopName)
    return nil
}

// ShopName Getter
func (r AlibabaEinvoiceBillSyncRequest) GetShopName() string {
    return r._shopName
}
// PayeeRegisterNo Setter
// 税务登记证号
func (r *AlibabaEinvoiceBillSyncRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceBillSyncRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// OrderId Setter
// 结算单订单ID
func (r *AlibabaEinvoiceBillSyncRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaEinvoiceBillSyncRequest) GetOrderId() string {
    return r._orderId
}
// SumPrice Setter
// 结算单总价格，小数点后2两位
func (r *AlibabaEinvoiceBillSyncRequest) SetSumPrice(_sumPrice string) error {
    r._sumPrice = _sumPrice
    r.Set("sum_price", _sumPrice)
    return nil
}

// SumPrice Getter
func (r AlibabaEinvoiceBillSyncRequest) GetSumPrice() string {
    return r._sumPrice
}
// Platform Setter
// 调用平台，用于区分同一个税号下多个店铺来源["TB:淘宝","ALIPAY:支付宝","TM:天猫","JD:京东","DD:当当","PP:拍拍","YX:易讯","EBAY:ebay","QQ:QQ网购","AMAZON:亚马逊","SN:苏宁","GM:国美","WPH:唯品会","JM:聚美","LF:乐蜂","MGJ:蘑菇街","JS:聚尚","PX:拍鞋","YT:银泰","YHD:1号店","VANCL:凡客","YL:邮乐","YG:优购","1688:阿里巴巴","POS:POS门店","ELEME:饿了么","OTHER:其他"]
func (r *AlibabaEinvoiceBillSyncRequest) SetPlatform(_platform string) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r AlibabaEinvoiceBillSyncRequest) GetPlatform() string {
    return r._platform
}
// Qrcode Setter
// 生成二维码参数，若不需要生成二维码，则不填
func (r *AlibabaEinvoiceBillSyncRequest) SetQrcode(_qrcode *QrCodeDo) error {
    r._qrcode = _qrcode
    r.Set("qrcode", _qrcode)
    return nil
}

// Qrcode Getter
func (r AlibabaEinvoiceBillSyncRequest) GetQrcode() *QrCodeDo {
    return r._qrcode
}
// BrandName Setter
// 品牌名称，不填默认=shop_name
func (r *AlibabaEinvoiceBillSyncRequest) SetBrandName(_brandName string) error {
    r._brandName = _brandName
    r.Set("brand_name", _brandName)
    return nil
}

// BrandName Getter
func (r AlibabaEinvoiceBillSyncRequest) GetBrandName() string {
    return r._brandName
}
// InvoicePrice Setter
// 结算单可开票总金额（不填=sumPrice），小数点后2两位
func (r *AlibabaEinvoiceBillSyncRequest) SetInvoicePrice(_invoicePrice string) error {
    r._invoicePrice = _invoicePrice
    r.Set("invoice_price", _invoicePrice)
    return nil
}

// InvoicePrice Getter
func (r AlibabaEinvoiceBillSyncRequest) GetInvoicePrice() string {
    return r._invoicePrice
}
// ShopPlatform Setter
// 开票店铺的平台，默认等于platform
func (r *AlibabaEinvoiceBillSyncRequest) SetShopPlatform(_shopPlatform string) error {
    r._shopPlatform = _shopPlatform
    r.Set("shop_platform", _shopPlatform)
    return nil
}

// ShopPlatform Getter
func (r AlibabaEinvoiceBillSyncRequest) GetShopPlatform() string {
    return r._shopPlatform
}
