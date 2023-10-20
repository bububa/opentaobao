package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotvpayorderqueryAPIRequest tv支付查询订单状态 API请求
// taobao.tvpay.order.query
//
// tv支付查询订单状态
type TaobaotvpayorderqueryAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 牌照方
	_license string
	// 业务订单号
	_bizOrderId string
	// 支付宝订单号
	_orderNo string
	// 订单类型
	_orderType string
	// 外部订单号
	_outOrderNo string
	// 是否淘系
	_isTao bool
}

// NewTaobaotvpayorderqueryRequest 初始化TaobaotvpayorderqueryAPIRequest对象
func NewTaobaotvpayorderqueryRequest() *TaobaotvpayorderqueryAPIRequest {
	return &TaobaotvpayorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotvpayorderqueryAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotvpayorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotvpayorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaotvpayorderqueryAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaotvpayorderqueryAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaotvpayorderqueryAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaotvpayorderqueryAPIRequest) GetFrom() string {
	return r._from
}

// SetLicense is License Setter
// 牌照方
func (r *TaobaotvpayorderqueryAPIRequest) SetLicense(_license string) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r TaobaotvpayorderqueryAPIRequest) GetLicense() string {
	return r._license
}

// SetBizOrderId is BizOrderId Setter
// 业务订单号
func (r *TaobaotvpayorderqueryAPIRequest) SetBizOrderId(_bizOrderId string) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaotvpayorderqueryAPIRequest) GetBizOrderId() string {
	return r._bizOrderId
}

// SetOrderNo is OrderNo Setter
// 支付宝订单号
func (r *TaobaotvpayorderqueryAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r TaobaotvpayorderqueryAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetOrderType is OrderType Setter
// 订单类型
func (r *TaobaotvpayorderqueryAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaotvpayorderqueryAPIRequest) GetOrderType() string {
	return r._orderType
}

// SetOutOrderNo is OutOrderNo Setter
// 外部订单号
func (r *TaobaotvpayorderqueryAPIRequest) SetOutOrderNo(_outOrderNo string) error {
	r._outOrderNo = _outOrderNo
	r.Set("out_order_no", _outOrderNo)
	return nil
}

// GetOutOrderNo OutOrderNo Getter
func (r TaobaotvpayorderqueryAPIRequest) GetOutOrderNo() string {
	return r._outOrderNo
}

// SetIsTao is IsTao Setter
// 是否淘系
func (r *TaobaotvpayorderqueryAPIRequest) SetIsTao(_isTao bool) error {
	r._isTao = _isTao
	r.Set("is_tao", _isTao)
	return nil
}

// GetIsTao IsTao Getter
func (r TaobaotvpayorderqueryAPIRequest) GetIsTao() bool {
	return r._isTao
}
