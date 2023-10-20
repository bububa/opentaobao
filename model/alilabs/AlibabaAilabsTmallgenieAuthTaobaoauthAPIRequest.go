package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthtaobaoauthAPIRequest 天猫精灵淘宝登录授权绑定接口 API请求
// alibaba.ailabs.tmallgenie.auth.taobaoauth
//
// 厂商获取用户淘宝授权之后，通过此接口获取天猫精灵授权，并绑定一台设备
type AlibabaailabstmallgenieauthtaobaoauthAPIRequest struct {
	model.Params
	// 授权信息
	_authParam *TopAuthReqDto
	// 设备信息
	_deviceParam *TopDeviceReqDto
}

// NewAlibabaailabstmallgenieauthtaobaoauthRequest 初始化AlibabaailabstmallgenieauthtaobaoauthAPIRequest对象
func NewAlibabaailabstmallgenieauthtaobaoauthRequest() *AlibabaailabstmallgenieauthtaobaoauthAPIRequest {
	return &AlibabaailabstmallgenieauthtaobaoauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthtaobaoauthAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.taobaoauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthtaobaoauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthtaobaoauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthParam is AuthParam Setter
// 授权信息
func (r *AlibabaailabstmallgenieauthtaobaoauthAPIRequest) SetAuthParam(_authParam *TopAuthReqDto) error {
	r._authParam = _authParam
	r.Set("auth_param", _authParam)
	return nil
}

// GetAuthParam AuthParam Getter
func (r AlibabaailabstmallgenieauthtaobaoauthAPIRequest) GetAuthParam() *TopAuthReqDto {
	return r._authParam
}

// SetDeviceParam is DeviceParam Setter
// 设备信息
func (r *AlibabaailabstmallgenieauthtaobaoauthAPIRequest) SetDeviceParam(_deviceParam *TopDeviceReqDto) error {
	r._deviceParam = _deviceParam
	r.Set("device_param", _deviceParam)
	return nil
}

// GetDeviceParam DeviceParam Getter
func (r AlibabaailabstmallgenieauthtaobaoauthAPIRequest) GetDeviceParam() *TopDeviceReqDto {
	return r._deviceParam
}
