package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthGetcodeAPIRequest 获取token API请求
// alibaba.ailabs.tmallgenie.auth.getcode
//
// 获取天猫精灵authCode
type AlibabaAilabsTmallgenieAuthGetcodeAPIRequest struct {
	model.Params
	// 授权参数
	_authParam *TopAuthReqDto
	// 设备参数
	_deviceParam *TopDeviceReqDto
}

// NewAlibabaAilabsTmallgenieAuthGetcodeRequest 初始化AlibabaAilabsTmallgenieAuthGetcodeAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthGetcodeRequest() *AlibabaAilabsTmallgenieAuthGetcodeAPIRequest {
	return &AlibabaAilabsTmallgenieAuthGetcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.getcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AuthParam Setter
// 授权参数
func (r *AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) SetAuthParam(_authParam *TopAuthReqDto) error {
	r._authParam = _authParam
	r.Set("auth_param", _authParam)
	return nil
}

// Get AuthParam Getter
func (r AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) GetAuthParam() *TopAuthReqDto {
	return r._authParam
}

// Set is DeviceParam Setter
// 设备参数
func (r *AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) SetDeviceParam(_deviceParam *TopDeviceReqDto) error {
	r._deviceParam = _deviceParam
	r.Set("device_param", _deviceParam)
	return nil
}

// Get DeviceParam Getter
func (r AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) GetDeviceParam() *TopDeviceReqDto {
	return r._deviceParam
}
