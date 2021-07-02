package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayOrderQueryAPIRequest tv支付查询订单状态 API请求
// taobao.tvpay.order.query
//
// tv支付查询订单状态
type TaobaoTvpayOrderQueryAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 业务订单号
	_bizOrderId string
	// 是否淘系
	_isTao bool
	// 支付宝订单号
	_orderNo string
	// 订单类型
	_orderType string
	// 外部订单号
	_outOrderNo string
	// 牌照方
	_license string
}

// NewTaobaoTvpayOrderQueryRequest 初始化TaobaoTvpayOrderQueryAPIRequest对象
func NewTaobaoTvpayOrderQueryRequest() *TaobaoTvpayOrderQueryAPIRequest {
	return &TaobaoTvpayOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayOrderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeviceId Setter
// 设备id
func (r *TaobaoTvpayOrderQueryAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// Get DeviceId Getter
func (r TaobaoTvpayOrderQueryAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// Set is From Setter
// 来源
func (r *TaobaoTvpayOrderQueryAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// Get From Getter
func (r TaobaoTvpayOrderQueryAPIRequest) GetFrom() string {
	return r._from
}

// Set is BizOrderId Setter
// 业务订单号
func (r *TaobaoTvpayOrderQueryAPIRequest) SetBizOrderId(_bizOrderId string) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// Get BizOrderId Getter
func (r TaobaoTvpayOrderQueryAPIRequest) GetBizOrderId() string {
	return r._bizOrderId
}

// Set is IsTao Setter
// 是否淘系
func (r *TaobaoTvpayOrderQueryAPIRequest) SetIsTao(_isTao bool) error {
	r._isTao = _isTao
	r.Set("is_tao", _isTao)
	return nil
}

// Get IsTao Getter
func (r TaobaoTvpayOrderQueryAPIRequest) GetIsTao() bool {
	return r._isTao
}

// Set is OrderNo Setter
// 支付宝订单号
func (r *TaobaoTvpayOrderQueryAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// Get OrderNo Getter
func (r TaobaoTvpayOrderQueryAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// Set is OrderType Setter
// 订单类型
func (r *TaobaoTvpayOrderQueryAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// Get OrderType Getter
func (r TaobaoTvpayOrderQueryAPIRequest) GetOrderType() string {
	return r._orderType
}

// Set is OutOrderNo Setter
// 外部订单号
func (r *TaobaoTvpayOrderQueryAPIRequest) SetOutOrderNo(_outOrderNo string) error {
	r._outOrderNo = _outOrderNo
	r.Set("out_order_no", _outOrderNo)
	return nil
}

// Get OutOrderNo Getter
func (r TaobaoTvpayOrderQueryAPIRequest) GetOutOrderNo() string {
	return r._outOrderNo
}

// Set is License Setter
// 牌照方
func (r *TaobaoTvpayOrderQueryAPIRequest) SetLicense(_license string) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// Get License Getter
func (r TaobaoTvpayOrderQueryAPIRequest) GetLicense() string {
	return r._license
}
