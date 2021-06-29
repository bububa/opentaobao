package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
结算单同步 APIRequest
alibaba.einvoice.bill.sync

电子发票业务，服务商同步结算单，包括结算单的增删改功能。最终用于开发票
*/
type AlibabaEinvoiceBillSyncRequest struct {
    model.Params

    // 结算商品单明细列表
    invoiceItems   []BillItemDo 

    // 结算单同步操作：=1插入，=2更新，=3废弃删除
    status   int64 

    // 结算单订单日期
    orderDate   string 

    // 店铺名称，与后台店铺名称保持一致
    shopName   string 

    // 税务登记证号
    payeeRegisterNo   string 

    // 结算单订单ID
    orderId   string 

    // 结算单总价格，小数点后2两位
    sumPrice   string 

    // 调用平台，用于区分同一个税号下多个店铺来源["TB:淘宝","ALIPAY:支付宝","TM:天猫","JD:京东","DD:当当","PP:拍拍","YX:易讯","EBAY:ebay","QQ:QQ网购","AMAZON:亚马逊","SN:苏宁","GM:国美","WPH:唯品会","JM:聚美","LF:乐蜂","MGJ:蘑菇街","JS:聚尚","PX:拍鞋","YT:银泰","YHD:1号店","VANCL:凡客","YL:邮乐","YG:优购","1688:阿里巴巴","POS:POS门店","ELEME:饿了么","OTHER:其他"]
    platform   string 

    // 生成二维码参数，若不需要生成二维码，则不填
    qrcode   *QrCodeDo 

    // 品牌名称，不填默认=shop_name
    brandName   string 

    // 结算单可开票总金额（不填=sumPrice），小数点后2两位
    invoicePrice   string 

    // 开票店铺的平台，默认等于platform
    shopPlatform   string 

}

func NewAlibabaEinvoiceBillSyncRequest() *AlibabaEinvoiceBillSyncRequest{
    return &AlibabaEinvoiceBillSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceBillSyncRequest) GetApiMethodName() string {
    return "alibaba.einvoice.bill.sync"
}

func (r AlibabaEinvoiceBillSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceBillSyncRequest) SetInvoiceItems(invoiceItems []BillItemDo) error {
    r.invoiceItems = invoiceItems
    r.Set("invoice_items", invoiceItems)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetInvoiceItems() []BillItemDo {
    return r.invoiceItems
}

func (r *AlibabaEinvoiceBillSyncRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetStatus() int64 {
    return r.status
}

func (r *AlibabaEinvoiceBillSyncRequest) SetOrderDate(orderDate string) error {
    r.orderDate = orderDate
    r.Set("order_date", orderDate)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetOrderDate() string {
    return r.orderDate
}

func (r *AlibabaEinvoiceBillSyncRequest) SetShopName(shopName string) error {
    r.shopName = shopName
    r.Set("shop_name", shopName)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetShopName() string {
    return r.shopName
}

func (r *AlibabaEinvoiceBillSyncRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoiceBillSyncRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaEinvoiceBillSyncRequest) SetSumPrice(sumPrice string) error {
    r.sumPrice = sumPrice
    r.Set("sum_price", sumPrice)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetSumPrice() string {
    return r.sumPrice
}

func (r *AlibabaEinvoiceBillSyncRequest) SetPlatform(platform string) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetPlatform() string {
    return r.platform
}

func (r *AlibabaEinvoiceBillSyncRequest) SetQrcode(qrcode *QrCodeDo) error {
    r.qrcode = qrcode
    r.Set("qrcode", qrcode)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetQrcode() *QrCodeDo {
    return r.qrcode
}

func (r *AlibabaEinvoiceBillSyncRequest) SetBrandName(brandName string) error {
    r.brandName = brandName
    r.Set("brand_name", brandName)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetBrandName() string {
    return r.brandName
}

func (r *AlibabaEinvoiceBillSyncRequest) SetInvoicePrice(invoicePrice string) error {
    r.invoicePrice = invoicePrice
    r.Set("invoice_price", invoicePrice)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetInvoicePrice() string {
    return r.invoicePrice
}

func (r *AlibabaEinvoiceBillSyncRequest) SetShopPlatform(shopPlatform string) error {
    r.shopPlatform = shopPlatform
    r.Set("shop_platform", shopPlatform)
    return nil
}

func (r AlibabaEinvoiceBillSyncRequest) GetShopPlatform() string {
    return r.shopPlatform
}

