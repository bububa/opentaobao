package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用加固接口 API请求
alibaba.security.jaq.app.shield

提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果
*/
type AlibabaSecurityJaqAppShieldRequest struct {
    model.Params
    // 待加固的应用信息
    _appInfo   *ScanAppInfo
    // 渠道列表,多渠道加固时填写
    _channel   *ShieldChannel
}

// 初始化AlibabaSecurityJaqAppShieldRequest对象
func NewAlibabaSecurityJaqAppShieldRequest() *AlibabaSecurityJaqAppShieldRequest{
    return &AlibabaSecurityJaqAppShieldRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppShieldRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.shield"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppShieldRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppInfo Setter
// 待加固的应用信息
func (r *AlibabaSecurityJaqAppShieldRequest) SetAppInfo(_appInfo *ScanAppInfo) error {
    r._appInfo = _appInfo
    r.Set("app_info", _appInfo)
    return nil
}

// AppInfo Getter
func (r AlibabaSecurityJaqAppShieldRequest) GetAppInfo() *ScanAppInfo {
    return r._appInfo
}
// Channel Setter
// 渠道列表,多渠道加固时填写
func (r *AlibabaSecurityJaqAppShieldRequest) SetChannel(_channel *ShieldChannel) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r AlibabaSecurityJaqAppShieldRequest) GetChannel() *ShieldChannel {
    return r._channel
}
