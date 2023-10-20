package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthgetcodeAPIRequest 获取token API请求
// alibaba.ailabs.tmallgenie.auth.getcode
//
// 获取天猫精灵authCode
type AlibabaailabstmallgenieauthgetcodeAPIRequest struct {
	model.Params
	// 授权参数
	_authParam *TopAuthReqDto
	// 设备参数
	_deviceParam *TopDeviceReqDto
}

// NewAlibabaailabstmallgenieauthgetcodeRequest 初始化AlibabaailabstmallgenieauthgetcodeAPIRequest对象
func NewAlibabaailabstmallgenieauthgetcodeRequest() *AlibabaailabstmallgenieauthgetcodeAPIRequest {
	return &AlibabaailabstmallgenieauthgetcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthgetcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.getcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthgetcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthgetcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthParam is AuthParam Setter
// 授权参数
func (r *AlibabaailabstmallgenieauthgetcodeAPIRequest) SetAuthParam(_authParam *TopAuthReqDto) error {
	r._authParam = _authParam
	r.Set("auth_param", _authParam)
	return nil
}

// GetAuthParam AuthParam Getter
func (r AlibabaailabstmallgenieauthgetcodeAPIRequest) GetAuthParam() *TopAuthReqDto {
	return r._authParam
}

// SetDeviceParam is DeviceParam Setter
// 设备参数
func (r *AlibabaailabstmallgenieauthgetcodeAPIRequest) SetDeviceParam(_deviceParam *TopDeviceReqDto) error {
	r._deviceParam = _deviceParam
	r.Set("device_param", _deviceParam)
	return nil
}

// GetDeviceParam DeviceParam Getter
func (r AlibabaailabstmallgenieauthgetcodeAPIRequest) GetDeviceParam() *TopDeviceReqDto {
	return r._deviceParam
}
