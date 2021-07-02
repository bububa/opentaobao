package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayAppinfoGetAPIRequest tv支付获取应用信息 API请求
// taobao.tvpay.appinfo.get
//
// tv支付获取应用信息
type TaobaoTvpayAppinfoGetAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 客户端版本号
	_clientVersion string
}

// NewTaobaoTvpayAppinfoGetRequest 初始化TaobaoTvpayAppinfoGetAPIRequest对象
func NewTaobaoTvpayAppinfoGetRequest() *TaobaoTvpayAppinfoGetAPIRequest {
	return &TaobaoTvpayAppinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAppinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.appinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAppinfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoTvpayAppinfoGetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoTvpayAppinfoGetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 来源
func (r *TaobaoTvpayAppinfoGetAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaoTvpayAppinfoGetAPIRequest) GetFrom() string {
	return r._from
}

// SetClientVersion is ClientVersion Setter
// 客户端版本号
func (r *TaobaoTvpayAppinfoGetAPIRequest) SetClientVersion(_clientVersion string) error {
	r._clientVersion = _clientVersion
	r.Set("client_version", _clientVersion)
	return nil
}

// GetClientVersion ClientVersion Getter
func (r TaobaoTvpayAppinfoGetAPIRequest) GetClientVersion() string {
	return r._clientVersion
}
