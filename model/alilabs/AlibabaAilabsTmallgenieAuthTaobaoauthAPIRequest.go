package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest
天猫精灵淘宝登录授权绑定接口 API请求
alibaba.ailabs.tmallgenie.auth.taobaoauth

厂商获取用户淘宝授权之后，通过此接口获取天猫精灵授权，并绑定一台设备 */
type AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest struct {
	model.Params
	// 授权信息
	_authParam *TopAuthReqDto
	// 设备信息
	_deviceParam *TopDeviceReqDto
}

// NewAlibabaAilabsTmallgenieAuthTaobaoauthRequest 初始化AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthTaobaoauthRequest() *AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest {
	return &AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.taobaoauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AuthParam Setter
// 授权信息
func (r *AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest) SetAuthParam(_authParam *TopAuthReqDto) error {
	r._authParam = _authParam
	r.Set("auth_param", _authParam)
	return nil
}

// Get AuthParam Getter
func (r AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest) GetAuthParam() *TopAuthReqDto {
	return r._authParam
}

// Set is DeviceParam Setter
// 设备信息
func (r *AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest) SetDeviceParam(_deviceParam *TopDeviceReqDto) error {
	r._deviceParam = _deviceParam
	r.Set("device_param", _deviceParam)
	return nil
}

// Get DeviceParam Getter
func (r AlibabaAilabsTmallgenieAuthTaobaoauthAPIRequest) GetDeviceParam() *TopDeviceReqDto {
	return r._deviceParam
}
