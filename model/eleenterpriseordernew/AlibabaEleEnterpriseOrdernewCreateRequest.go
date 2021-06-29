package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 APIRequest
alibaba.ele.enterprise.ordernew.create

创建订单
*/
type AlibabaEleEnterpriseOrdernewCreateRequest struct {
    model.Params

    // 订单来源地址经度
    longitude   string 

    // 餐厅Id
    erestaurantId   string 

    // 使用的券号
    couponSn   string 

    // 订单备注信息
    description   string 

    // 电话号码，主要号码必须是手机号；多个手机号以逗号分隔
    phones   string 

    // 订单来源IP地址
    ip   string 

    // 订单来源地址纬度
    latitude   string 

    // 购物车Id（创建购物车返回的购物车id）
    cartId   string 

    // 第三方订单Id（需保证唯一）
    tpOrderId   string 

    // 送餐地址
    address   string 

    // 收餐人姓名
    consignee   string 

    // 暂时不用传（忽略此字段）
    deliverTime   string 

    // 纳税人识别号
    invoiceNumber   string 

    // 发票抬头（个人发票请填写个人），不传表示不要发票
    invoice   string 

    // 发票类型（发票类型, 1: 个人, 2: 企业; 空为兼容数据, 由商户判断发票类型）
    invoiceType   int64 

}

func NewAlibabaEleEnterpriseOrdernewCreateRequest() *AlibabaEleEnterpriseOrdernewCreateRequest{
    return &AlibabaEleEnterpriseOrdernewCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.create"
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetLongitude() string {
    return r.longitude
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetErestaurantId(erestaurantId string) error {
    r.erestaurantId = erestaurantId
    r.Set("erestaurant_id", erestaurantId)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetErestaurantId() string {
    return r.erestaurantId
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetCouponSn(couponSn string) error {
    r.couponSn = couponSn
    r.Set("coupon_sn", couponSn)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetCouponSn() string {
    return r.couponSn
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetDescription() string {
    return r.description
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetPhones(phones string) error {
    r.phones = phones
    r.Set("phones", phones)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetPhones() string {
    return r.phones
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetIp() string {
    return r.ip
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetLatitude() string {
    return r.latitude
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetCartId(cartId string) error {
    r.cartId = cartId
    r.Set("cart_id", cartId)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetCartId() string {
    return r.cartId
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetTpOrderId(tpOrderId string) error {
    r.tpOrderId = tpOrderId
    r.Set("tp_order_id", tpOrderId)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetTpOrderId() string {
    return r.tpOrderId
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetAddress() string {
    return r.address
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetConsignee(consignee string) error {
    r.consignee = consignee
    r.Set("consignee", consignee)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetConsignee() string {
    return r.consignee
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetDeliverTime(deliverTime string) error {
    r.deliverTime = deliverTime
    r.Set("deliver_time", deliverTime)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetDeliverTime() string {
    return r.deliverTime
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetInvoiceNumber(invoiceNumber string) error {
    r.invoiceNumber = invoiceNumber
    r.Set("invoice_number", invoiceNumber)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetInvoiceNumber() string {
    return r.invoiceNumber
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetInvoice(invoice string) error {
    r.invoice = invoice
    r.Set("invoice", invoice)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetInvoice() string {
    return r.invoice
}

func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetInvoiceType(invoiceType int64) error {
    r.invoiceType = invoiceType
    r.Set("invoice_type", invoiceType)
    return nil
}

func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetInvoiceType() int64 {
    return r.invoiceType
}

