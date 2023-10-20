package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotvpayaccessdatagetAPIRequest tv支付 API请求
// taobao.tvpay.access.data.get
//
// 在匿名用户支付后尝试为其登陆绑定的淘宝账号
type TaobaotvpayaccessdatagetAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 账号客户端版本
	_accountClientVersion string
	// 订单id
	_outOrderNo string
}

// NewTaobaotvpayaccessdatagetRequest 初始化TaobaotvpayaccessdatagetAPIRequest对象
func NewTaobaotvpayaccessdatagetRequest() *TaobaotvpayaccessdatagetAPIRequest {
	return &TaobaotvpayaccessdatagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotvpayaccessdatagetAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.access.data.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotvpayaccessdatagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotvpayaccessdatagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaotvpayaccessdatagetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaotvpayaccessdatagetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaotvpayaccessdatagetAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaotvpayaccessdatagetAPIRequest) GetFrom() string {
	return r._from
}

// SetAccountClientVersion is AccountClientVersion Setter
// 账号客户端版本
func (r *TaobaotvpayaccessdatagetAPIRequest) SetAccountClientVersion(_accountClientVersion string) error {
	r._accountClientVersion = _accountClientVersion
	r.Set("account_client_version", _accountClientVersion)
	return nil
}

// GetAccountClientVersion AccountClientVersion Getter
func (r TaobaotvpayaccessdatagetAPIRequest) GetAccountClientVersion() string {
	return r._accountClientVersion
}

// SetOutOrderNo is OutOrderNo Setter
// 订单id
func (r *TaobaotvpayaccessdatagetAPIRequest) SetOutOrderNo(_outOrderNo string) error {
	r._outOrderNo = _outOrderNo
	r.Set("out_order_no", _outOrderNo)
	return nil
}

// GetOutOrderNo OutOrderNo Getter
func (r TaobaotvpayaccessdatagetAPIRequest) GetOutOrderNo() string {
	return r._outOrderNo
}
