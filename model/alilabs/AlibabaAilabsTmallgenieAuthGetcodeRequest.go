package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取token APIRequest
alibaba.ailabs.tmallgenie.auth.getcode

获取天猫精灵authCode
*/
type AlibabaAilabsTmallgenieAuthGetcodeRequest struct {
    model.Params

    // 授权参数
    authParam   *TopAuthReqDto 

    // 设备参数
    deviceParam   *TopDeviceReqDto 

}

func NewAlibabaAilabsTmallgenieAuthGetcodeRequest() *AlibabaAilabsTmallgenieAuthGetcodeRequest{
    return &AlibabaAilabsTmallgenieAuthGetcodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthGetcodeRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.getcode"
}

func (r AlibabaAilabsTmallgenieAuthGetcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthGetcodeRequest) SetAuthParam(authParam *TopAuthReqDto) error {
    r.authParam = authParam
    r.Set("auth_param", authParam)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthGetcodeRequest) GetAuthParam() *TopAuthReqDto {
    return r.authParam
}

func (r *AlibabaAilabsTmallgenieAuthGetcodeRequest) SetDeviceParam(deviceParam *TopDeviceReqDto) error {
    r.deviceParam = deviceParam
    r.Set("device_param", deviceParam)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthGetcodeRequest) GetDeviceParam() *TopDeviceReqDto {
    return r.deviceParam
}

