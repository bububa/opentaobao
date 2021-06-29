package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取token API请求
alibaba.ailabs.tmallgenie.auth.getcode

获取天猫精灵authCode
*/
type AlibabaAilabsTmallgenieAuthGetcodeRequest struct {
    model.Params
    // 授权参数
    _authParam   *TopAuthReqDto
    // 设备参数
    _deviceParam   *TopDeviceReqDto
}

// 初始化AlibabaAilabsTmallgenieAuthGetcodeRequest对象
func NewAlibabaAilabsTmallgenieAuthGetcodeRequest() *AlibabaAilabsTmallgenieAuthGetcodeRequest{
    return &AlibabaAilabsTmallgenieAuthGetcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthGetcodeRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.getcode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthGetcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AuthParam Setter
// 授权参数
func (r *AlibabaAilabsTmallgenieAuthGetcodeRequest) SetAuthParam(_authParam *TopAuthReqDto) error {
    r._authParam = _authParam
    r.Set("auth_param", _authParam)
    return nil
}

// AuthParam Getter
func (r AlibabaAilabsTmallgenieAuthGetcodeRequest) GetAuthParam() *TopAuthReqDto {
    return r._authParam
}
// DeviceParam Setter
// 设备参数
func (r *AlibabaAilabsTmallgenieAuthGetcodeRequest) SetDeviceParam(_deviceParam *TopDeviceReqDto) error {
    r._deviceParam = _deviceParam
    r.Set("device_param", _deviceParam)
    return nil
}

// DeviceParam Getter
func (r AlibabaAilabsTmallgenieAuthGetcodeRequest) GetDeviceParam() *TopDeviceReqDto {
    return r._deviceParam
}
