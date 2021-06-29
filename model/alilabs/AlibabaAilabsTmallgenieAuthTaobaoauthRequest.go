package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵淘宝登录授权绑定接口 APIRequest
alibaba.ailabs.tmallgenie.auth.taobaoauth

厂商获取用户淘宝授权之后，通过此接口获取天猫精灵授权，并绑定一台设备
*/
type AlibabaAilabsTmallgenieAuthTaobaoauthRequest struct {
    model.Params

    // 授权信息
    authParam   *TopAuthReqDto 

    // 设备信息
    deviceParam   *TopDeviceReqDto 

}

func NewAlibabaAilabsTmallgenieAuthTaobaoauthRequest() *AlibabaAilabsTmallgenieAuthTaobaoauthRequest{
    return &AlibabaAilabsTmallgenieAuthTaobaoauthRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthTaobaoauthRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.taobaoauth"
}

func (r AlibabaAilabsTmallgenieAuthTaobaoauthRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthTaobaoauthRequest) SetAuthParam(authParam *TopAuthReqDto) error {
    r.authParam = authParam
    r.Set("auth_param", authParam)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthTaobaoauthRequest) GetAuthParam() *TopAuthReqDto {
    return r.authParam
}

func (r *AlibabaAilabsTmallgenieAuthTaobaoauthRequest) SetDeviceParam(deviceParam *TopDeviceReqDto) error {
    r.deviceParam = deviceParam
    r.Set("device_param", deviceParam)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthTaobaoauthRequest) GetDeviceParam() *TopDeviceReqDto {
    return r.deviceParam
}

