package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotvpayorderprecreateAPIRequest tv支付预下单 API请求
// taobao.tvpay.order.precreate
//
// tv支付预下单
type TaobaotvpayorderprecreateAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 牌照方
	_license string
	// 订单详情
	_data string
}

// NewTaobaotvpayorderprecreateRequest 初始化TaobaotvpayorderprecreateAPIRequest对象
func NewTaobaotvpayorderprecreateRequest() *TaobaotvpayorderprecreateAPIRequest {
	return &TaobaotvpayorderprecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotvpayorderprecreateAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.order.precreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotvpayorderprecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotvpayorderprecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaotvpayorderprecreateAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaotvpayorderprecreateAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaotvpayorderprecreateAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaotvpayorderprecreateAPIRequest) GetFrom() string {
	return r._from
}

// SetLicense is License Setter
// 牌照方
func (r *TaobaotvpayorderprecreateAPIRequest) SetLicense(_license string) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r TaobaotvpayorderprecreateAPIRequest) GetLicense() string {
	return r._license
}

// SetData is Data Setter
// 订单详情
func (r *TaobaotvpayorderprecreateAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaotvpayorderprecreateAPIRequest) GetData() string {
	return r._data
}
