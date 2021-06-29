package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取资源文件 API请求
alibaba.security.jaq.resource.fetch

在前向化验证流程中提供资源文件服务
*/
type AlibabaSecurityJaqResourceFetchRequest struct {
    model.Params
    // 设备类型可能值有：android ios wp
    deviceType   string
    // 分辨率
    dpi   string
    // 语言类型 zh_CN en_US
    lang   string
}

// 初始化AlibabaSecurityJaqResourceFetchRequest对象
func NewAlibabaSecurityJaqResourceFetchRequest() *AlibabaSecurityJaqResourceFetchRequest{
    return &AlibabaSecurityJaqResourceFetchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqResourceFetchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.resource.fetch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqResourceFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceType Setter
// 设备类型可能值有：android ios wp
func (r *AlibabaSecurityJaqResourceFetchRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaSecurityJaqResourceFetchRequest) GetDeviceType() string {
    return r.deviceType
}
// Dpi Setter
// 分辨率
func (r *AlibabaSecurityJaqResourceFetchRequest) SetDpi(dpi string) error {
    r.dpi = dpi
    r.Set("dpi", dpi)
    return nil
}

// Dpi Getter
func (r AlibabaSecurityJaqResourceFetchRequest) GetDpi() string {
    return r.dpi
}
// Lang Setter
// 语言类型 zh_CN en_US
func (r *AlibabaSecurityJaqResourceFetchRequest) SetLang(lang string) error {
    r.lang = lang
    r.Set("lang", lang)
    return nil
}

// Lang Getter
func (r AlibabaSecurityJaqResourceFetchRequest) GetLang() string {
    return r.lang
}
