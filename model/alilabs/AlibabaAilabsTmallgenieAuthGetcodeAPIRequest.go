package alilabs

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) Reset() {
	r._authParam = nil
	r._deviceParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.getcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthParam is AuthParam Setter
// 授权参数
func (r *AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) SetAuthParam(_authParam *TopAuthReqDto) error {
	r._authParam = _authParam
	r.Set("auth_param", _authParam)
	return nil
}

// GetAuthParam AuthParam Getter
func (r AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) GetAuthParam() *TopAuthReqDto {
	return r._authParam
}

// SetDeviceParam is DeviceParam Setter
// 设备参数
func (r *AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) SetDeviceParam(_deviceParam *TopDeviceReqDto) error {
	r._deviceParam = _deviceParam
	r.Set("device_param", _deviceParam)
	return nil
}

// GetDeviceParam DeviceParam Getter
func (r AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) GetDeviceParam() *TopDeviceReqDto {
	return r._deviceParam
}

var poolAlibabaAilabsTmallgenieAuthGetcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsTmallgenieAuthGetcodeRequest()
	},
}

// GetAlibabaAilabsTmallgenieAuthGetcodeRequest 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthGetcodeAPIRequest
func GetAlibabaAilabsTmallgenieAuthGetcodeAPIRequest() *AlibabaAilabsTmallgenieAuthGetcodeAPIRequest {
	return poolAlibabaAilabsTmallgenieAuthGetcodeAPIRequest.Get().(*AlibabaAilabsTmallgenieAuthGetcodeAPIRequest)
}

// ReleaseAlibabaAilabsTmallgenieAuthGetcodeAPIRequest 将 AlibabaAilabsTmallgenieAuthGetcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthGetcodeAPIRequest(v *AlibabaAilabsTmallgenieAuthGetcodeAPIRequest) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthGetcodeAPIRequest.Put(v)
}
