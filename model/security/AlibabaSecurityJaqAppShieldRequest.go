package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用加固接口 APIRequest
alibaba.security.jaq.app.shield

提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果
*/
type AlibabaSecurityJaqAppShieldRequest struct {
    model.Params

    // 待加固的应用信息
    appInfo   *ScanAppInfo 

    // 渠道列表,多渠道加固时填写
    channel   *ShieldChannel 

}

func NewAlibabaSecurityJaqAppShieldRequest() *AlibabaSecurityJaqAppShieldRequest{
    return &AlibabaSecurityJaqAppShieldRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqAppShieldRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.shield"
}

func (r AlibabaSecurityJaqAppShieldRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqAppShieldRequest) SetAppInfo(appInfo *ScanAppInfo) error {
    r.appInfo = appInfo
    r.Set("app_info", appInfo)
    return nil
}

func (r AlibabaSecurityJaqAppShieldRequest) GetAppInfo() *ScanAppInfo {
    return r.appInfo
}

func (r *AlibabaSecurityJaqAppShieldRequest) SetChannel(channel *ShieldChannel) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r AlibabaSecurityJaqAppShieldRequest) GetChannel() *ShieldChannel {
    return r.channel
}

