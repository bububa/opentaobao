package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewCreateAPIRequest 创建订单 API请求
// alibaba.ele.enterprise.ordernew.create
//
// 创建订单
type AlibabaEleEnterpriseOrdernewCreateAPIRequest struct {
	model.Params
	// 订单来源地址经度
	_longitude string
	// 餐厅Id
	_erestaurantId string
	// 使用的券号
	_couponSn string
	// 订单备注信息
	_description string
	// 电话号码，主要号码必须是手机号；多个手机号以逗号分隔
	_phones string
	// 订单来源IP地址
	_ip string
	// 订单来源地址纬度
	_latitude string
	// 购物车Id（创建购物车返回的购物车id）
	_cartId string
	// 第三方订单Id（需保证唯一）
	_tpOrderId string
	// 送餐地址
	_address string
	// 收餐人姓名
	_consignee string
	// 暂时不用传（忽略此字段）
	_deliverTime string
	// 纳税人识别号
	_invoiceNumber string
	// 发票抬头（个人发票请填写个人），不传表示不要发票
	_invoice string
	// 发票类型（发票类型, 1: 个人, 2: 企业; 空为兼容数据, 由商户判断发票类型）
	_invoiceType int64
}

// NewAlibabaEleEnterpriseOrdernewCreateRequest 初始化AlibabaEleEnterpriseOrdernewCreateAPIRequest对象
func NewAlibabaEleEnterpriseOrdernewCreateRequest() *AlibabaEleEnterpriseOrdernewCreateAPIRequest {
	return &AlibabaEleEnterpriseOrdernewCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLongitude is Longitude Setter
// 订单来源地址经度
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetErestaurantId is ErestaurantId Setter
// 餐厅Id
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetErestaurantId(_erestaurantId string) error {
	r._erestaurantId = _erestaurantId
	r.Set("erestaurant_id", _erestaurantId)
	return nil
}

// GetErestaurantId ErestaurantId Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetErestaurantId() string {
	return r._erestaurantId
}

// SetCouponSn is CouponSn Setter
// 使用的券号
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetCouponSn(_couponSn string) error {
	r._couponSn = _couponSn
	r.Set("coupon_sn", _couponSn)
	return nil
}

// GetCouponSn CouponSn Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetCouponSn() string {
	return r._couponSn
}

// SetDescription is Description Setter
// 订单备注信息
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetDescription() string {
	return r._description
}

// SetPhones is Phones Setter
// 电话号码，主要号码必须是手机号；多个手机号以逗号分隔
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetPhones(_phones string) error {
	r._phones = _phones
	r.Set("phones", _phones)
	return nil
}

// GetPhones Phones Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetPhones() string {
	return r._phones
}

// SetIp is Ip Setter
// 订单来源IP地址
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetIp() string {
	return r._ip
}

// SetLatitude is Latitude Setter
// 订单来源地址纬度
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetCartId is CartId Setter
// 购物车Id（创建购物车返回的购物车id）
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetCartId(_cartId string) error {
	r._cartId = _cartId
	r.Set("cart_id", _cartId)
	return nil
}

// GetCartId CartId Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetCartId() string {
	return r._cartId
}

// SetTpOrderId is TpOrderId Setter
// 第三方订单Id（需保证唯一）
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetTpOrderId(_tpOrderId string) error {
	r._tpOrderId = _tpOrderId
	r.Set("tp_order_id", _tpOrderId)
	return nil
}

// GetTpOrderId TpOrderId Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetTpOrderId() string {
	return r._tpOrderId
}

// SetAddress is Address Setter
// 送餐地址
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetAddress() string {
	return r._address
}

// SetConsignee is Consignee Setter
// 收餐人姓名
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetConsignee(_consignee string) error {
	r._consignee = _consignee
	r.Set("consignee", _consignee)
	return nil
}

// GetConsignee Consignee Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetConsignee() string {
	return r._consignee
}

// SetDeliverTime is DeliverTime Setter
// 暂时不用传（忽略此字段）
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetDeliverTime(_deliverTime string) error {
	r._deliverTime = _deliverTime
	r.Set("deliver_time", _deliverTime)
	return nil
}

// GetDeliverTime DeliverTime Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetDeliverTime() string {
	return r._deliverTime
}

// SetInvoiceNumber is InvoiceNumber Setter
// 纳税人识别号
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetInvoiceNumber(_invoiceNumber string) error {
	r._invoiceNumber = _invoiceNumber
	r.Set("invoice_number", _invoiceNumber)
	return nil
}

// GetInvoiceNumber InvoiceNumber Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetInvoiceNumber() string {
	return r._invoiceNumber
}

// SetInvoice is Invoice Setter
// 发票抬头（个人发票请填写个人），不传表示不要发票
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetInvoice(_invoice string) error {
	r._invoice = _invoice
	r.Set("invoice", _invoice)
	return nil
}

// GetInvoice Invoice Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetInvoice() string {
	return r._invoice
}

// SetInvoiceType is InvoiceType Setter
// 发票类型（发票类型, 1: 个人, 2: 企业; 空为兼容数据, 由商户判断发票类型）
func (r *AlibabaEleEnterpriseOrdernewCreateAPIRequest) SetInvoiceType(_invoiceType int64) error {
	r._invoiceType = _invoiceType
	r.Set("invoice_type", _invoiceType)
	return nil
}

// GetInvoiceType InvoiceType Getter
func (r AlibabaEleEnterpriseOrdernewCreateAPIRequest) GetInvoiceType() int64 {
	return r._invoiceType
}
