package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsOrderShengxianConfirmAPIRequest
物流宝生鲜冷链的发货 API请求
taobao.logistics.order.shengxian.confirm

优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。 */
type TaobaoLogisticsOrderShengxianConfirmAPIRequest struct {
	model.Params
	// 淘宝交易ID
	_tid int64
	// 运单号.具体一个物流公司的真实运单号码。支持多个运单号，多个运单号之间用英文分号（;）隔开，不能重复。淘宝官方物流会校验，请谨慎传入；
	_outSid string
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font>如果service_code不为空，默认使用service_code如果service_code为空，sender_id不为空，使用sender_id对应的地址发货如果service_code与sender_id都为空，使用默认地址发货
	_senderId int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br>
	_cancelId int64
	// 商家的IP地址
	_sellerIp string
	// 物流订单ID 。同淘宝交易订单互斥，淘宝交易号存在，，以淘宝交易号为准
	_logisId int64
	// 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)如果service_code为空，默认通过sender_id来发货
	_serviceCode string
	// 1：冷链。0：常温
	_deliveryType int64
}

// NewTaobaoLogisticsOrderShengxianConfirmRequest 初始化TaobaoLogisticsOrderShengxianConfirmAPIRequest对象
func NewTaobaoLogisticsOrderShengxianConfirmRequest() *TaobaoLogisticsOrderShengxianConfirmAPIRequest {
	return &TaobaoLogisticsOrderShengxianConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOrderShengxianConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.order.shengxian.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOrderShengxianConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsOrderShengxianConfirmAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoLogisticsOrderShengxianConfirmAPIRequest) GetTid() int64 {
	return r._tid
}

// Set is OutSid Setter
// 运单号.具体一个物流公司的真实运单号码。支持多个运单号，多个运单号之间用英文分号（;）隔开，不能重复。淘宝官方物流会校验，请谨慎传入；
func (r *TaobaoLogisticsOrderShengxianConfirmAPIRequest) SetOutSid(_outSid string) error {
	r._outSid = _outSid
	r.Set("out_sid", _outSid)
	return nil
}

// Get OutSid Getter
func (r TaobaoLogisticsOrderShengxianConfirmAPIRequest) GetOutSid() string {
	return r._outSid
}

// Set is SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font>如果service_code不为空，默认使用service_code如果service_code为空，sender_id不为空，使用sender_id对应的地址发货如果service_code与sender_id都为空，使用默认地址发货
func (r *TaobaoLogisticsOrderShengxianConfirmAPIRequest) SetSenderId(_senderId int64) error {
	r._senderId = _senderId
	r.Set("sender_id", _senderId)
	return nil
}

// Get SenderId Getter
func (r TaobaoLogisticsOrderShengxianConfirmAPIRequest) GetSenderId() int64 {
	return r._senderId
}

// Set is CancelId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br>
func (r *TaobaoLogisticsOrderShengxianConfirmAPIRequest) SetCancelId(_cancelId int64) error {
	r._cancelId = _cancelId
	r.Set("cancel_id", _cancelId)
	return nil
}

// Get CancelId Getter
func (r TaobaoLogisticsOrderShengxianConfirmAPIRequest) GetCancelId() int64 {
	return r._cancelId
}

// Set is SellerIp Setter
// 商家的IP地址
func (r *TaobaoLogisticsOrderShengxianConfirmAPIRequest) SetSellerIp(_sellerIp string) error {
	r._sellerIp = _sellerIp
	r.Set("seller_ip", _sellerIp)
	return nil
}

// Get SellerIp Getter
func (r TaobaoLogisticsOrderShengxianConfirmAPIRequest) GetSellerIp() string {
	return r._sellerIp
}

// Set is LogisId Setter
// 物流订单ID 。同淘宝交易订单互斥，淘宝交易号存在，，以淘宝交易号为准
func (r *TaobaoLogisticsOrderShengxianConfirmAPIRequest) SetLogisId(_logisId int64) error {
	r._logisId = _logisId
	r.Set("logis_id", _logisId)
	return nil
}

// Get LogisId Getter
func (r TaobaoLogisticsOrderShengxianConfirmAPIRequest) GetLogisId() int64 {
	return r._logisId
}

// Set is ServiceCode Setter
// 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)如果service_code为空，默认通过sender_id来发货
func (r *TaobaoLogisticsOrderShengxianConfirmAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// Get ServiceCode Getter
func (r TaobaoLogisticsOrderShengxianConfirmAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

// Set is DeliveryType Setter
// 1：冷链。0：常温
func (r *TaobaoLogisticsOrderShengxianConfirmAPIRequest) SetDeliveryType(_deliveryType int64) error {
	r._deliveryType = _deliveryType
	r.Set("delivery_type", _deliveryType)
	return nil
}

// Get DeliveryType Getter
func (r TaobaoLogisticsOrderShengxianConfirmAPIRequest) GetDeliveryType() int64 {
	return r._deliveryType
}
