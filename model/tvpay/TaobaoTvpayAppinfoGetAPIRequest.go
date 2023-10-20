package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotvpayappinfogetAPIRequest tv支付获取应用信息 API请求
// taobao.tvpay.appinfo.get
//
// tv支付获取应用信息
type TaobaotvpayappinfogetAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 客户端版本号
	_clientVersion string
}

// NewTaobaotvpayappinfogetRequest 初始化TaobaotvpayappinfogetAPIRequest对象
func NewTaobaotvpayappinfogetRequest() *TaobaotvpayappinfogetAPIRequest {
	return &TaobaotvpayappinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotvpayappinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.appinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotvpayappinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotvpayappinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaotvpayappinfogetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaotvpayappinfogetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaotvpayappinfogetAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaotvpayappinfogetAPIRequest) GetFrom() string {
	return r._from
}

// SetClientVersion is ClientVersion Setter
// 客户端版本号
func (r *TaobaotvpayappinfogetAPIRequest) SetClientVersion(_clientVersion string) error {
	r._clientVersion = _clientVersion
	r.Set("client_version", _clientVersion)
	return nil
}

// GetClientVersion ClientVersion Getter
func (r TaobaotvpayappinfogetAPIRequest) GetClientVersion() string {
	return r._clientVersion
}
