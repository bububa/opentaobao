package eleenterpriseordernew

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 API请求
alibaba.ele.enterprise.ordernew.create

创建订单
*/
type AlibabaEleEnterpriseOrdernewCreateRequest struct {
    model.Params
    // 订单来源地址经度
    _longitude   string
    // 餐厅Id
    _erestaurantId   string
    // 使用的券号
    _couponSn   string
    // 订单备注信息
    _description   string
    // 电话号码，主要号码必须是手机号；多个手机号以逗号分隔
    _phones   string
    // 订单来源IP地址
    _ip   string
    // 订单来源地址纬度
    _latitude   string
    // 购物车Id（创建购物车返回的购物车id）
    _cartId   string
    // 第三方订单Id（需保证唯一）
    _tpOrderId   string
    // 送餐地址
    _address   string
    // 收餐人姓名
    _consignee   string
    // 暂时不用传（忽略此字段）
    _deliverTime   string
    // 纳税人识别号
    _invoiceNumber   string
    // 发票抬头（个人发票请填写个人），不传表示不要发票
    _invoice   string
    // 发票类型（发票类型, 1: 个人, 2: 企业; 空为兼容数据, 由商户判断发票类型）
    _invoiceType   int64
}

// 初始化AlibabaEleEnterpriseOrdernewCreateRequest对象
func NewAlibabaEleEnterpriseOrdernewCreateRequest() *AlibabaEleEnterpriseOrdernewCreateRequest{
    return &AlibabaEleEnterpriseOrdernewCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetApiMethodName() string {
    return "alibaba.ele.enterprise.ordernew.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Longitude Setter
// 订单来源地址经度
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetLongitude() string {
    return r._longitude
}
// ErestaurantId Setter
// 餐厅Id
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetErestaurantId(_erestaurantId string) error {
    r._erestaurantId = _erestaurantId
    r.Set("erestaurant_id", _erestaurantId)
    return nil
}

// ErestaurantId Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetErestaurantId() string {
    return r._erestaurantId
}
// CouponSn Setter
// 使用的券号
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetCouponSn(_couponSn string) error {
    r._couponSn = _couponSn
    r.Set("coupon_sn", _couponSn)
    return nil
}

// CouponSn Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetCouponSn() string {
    return r._couponSn
}
// Description Setter
// 订单备注信息
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetDescription() string {
    return r._description
}
// Phones Setter
// 电话号码，主要号码必须是手机号；多个手机号以逗号分隔
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetPhones(_phones string) error {
    r._phones = _phones
    r.Set("phones", _phones)
    return nil
}

// Phones Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetPhones() string {
    return r._phones
}
// Ip Setter
// 订单来源IP地址
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetIp() string {
    return r._ip
}
// Latitude Setter
// 订单来源地址纬度
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetLatitude() string {
    return r._latitude
}
// CartId Setter
// 购物车Id（创建购物车返回的购物车id）
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetCartId(_cartId string) error {
    r._cartId = _cartId
    r.Set("cart_id", _cartId)
    return nil
}

// CartId Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetCartId() string {
    return r._cartId
}
// TpOrderId Setter
// 第三方订单Id（需保证唯一）
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetTpOrderId(_tpOrderId string) error {
    r._tpOrderId = _tpOrderId
    r.Set("tp_order_id", _tpOrderId)
    return nil
}

// TpOrderId Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetTpOrderId() string {
    return r._tpOrderId
}
// Address Setter
// 送餐地址
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetAddress() string {
    return r._address
}
// Consignee Setter
// 收餐人姓名
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetConsignee(_consignee string) error {
    r._consignee = _consignee
    r.Set("consignee", _consignee)
    return nil
}

// Consignee Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetConsignee() string {
    return r._consignee
}
// DeliverTime Setter
// 暂时不用传（忽略此字段）
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetDeliverTime(_deliverTime string) error {
    r._deliverTime = _deliverTime
    r.Set("deliver_time", _deliverTime)
    return nil
}

// DeliverTime Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetDeliverTime() string {
    return r._deliverTime
}
// InvoiceNumber Setter
// 纳税人识别号
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetInvoiceNumber(_invoiceNumber string) error {
    r._invoiceNumber = _invoiceNumber
    r.Set("invoice_number", _invoiceNumber)
    return nil
}

// InvoiceNumber Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetInvoiceNumber() string {
    return r._invoiceNumber
}
// Invoice Setter
// 发票抬头（个人发票请填写个人），不传表示不要发票
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetInvoice(_invoice string) error {
    r._invoice = _invoice
    r.Set("invoice", _invoice)
    return nil
}

// Invoice Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetInvoice() string {
    return r._invoice
}
// InvoiceType Setter
// 发票类型（发票类型, 1: 个人, 2: 企业; 空为兼容数据, 由商户判断发票类型）
func (r *AlibabaEleEnterpriseOrdernewCreateRequest) SetInvoiceType(_invoiceType int64) error {
    r._invoiceType = _invoiceType
    r.Set("invoice_type", _invoiceType)
    return nil
}

// InvoiceType Getter
func (r AlibabaEleEnterpriseOrdernewCreateRequest) GetInvoiceType() int64 {
    return r._invoiceType
}
