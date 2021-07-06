package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayAuthQueryAPIRequest tv支付授权查询 API请求
// taobao.tvpay.auth.query
//
// 查询该用户在指定设备上是否有支付授权
type TaobaoTvpayAuthQueryAPIRequest struct {
	model.Params
	// 设备号
	_deviceId string
	// 来源
	_from string
	// 业务订单号
	_bizOrderId string
	// 支付宝订单号
	_orderNo string
	// 外部订单号
	_outOrderNo string
	// 是否淘系
	_isTao bool
}

// NewTaobaoTvpayAuthQueryRequest 初始化TaobaoTvpayAuthQueryAPIRequest对象
func NewTaobaoTvpayAuthQueryRequest() *TaobaoTvpayAuthQueryAPIRequest {
	return &TaobaoTvpayAuthQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAuthQueryAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.auth.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAuthQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeviceId is DeviceId Setter
// 设备号
func (r *TaobaoTvpayAuthQueryAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaoTvpayAuthQueryAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetFrom() string {
	return r._from
}

// SetBizOrderId is BizOrderId Setter
// 业务订单号
func (r *TaobaoTvpayAuthQueryAPIRequest) SetBizOrderId(_bizOrderId string) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetBizOrderId() string {
	return r._bizOrderId
}

// SetOrderNo is OrderNo Setter
// 支付宝订单号
func (r *TaobaoTvpayAuthQueryAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetOutOrderNo is OutOrderNo Setter
// 外部订单号
func (r *TaobaoTvpayAuthQueryAPIRequest) SetOutOrderNo(_outOrderNo string) error {
	r._outOrderNo = _outOrderNo
	r.Set("out_order_no", _outOrderNo)
	return nil
}

// GetOutOrderNo OutOrderNo Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetOutOrderNo() string {
	return r._outOrderNo
}

// SetIsTao is IsTao Setter
// 是否淘系
func (r *TaobaoTvpayAuthQueryAPIRequest) SetIsTao(_isTao bool) error {
	r._isTao = _isTao
	r.Set("is_tao", _isTao)
	return nil
}

// GetIsTao IsTao Getter
func (r TaobaoTvpayAuthQueryAPIRequest) GetIsTao() bool {
	return r._isTao
}
