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
type AlibabaSecurityJaqResourceFetchAPIRequest struct {
    model.Params
    // 设备类型可能值有：android ios wp
    _deviceType   string
    // 分辨率
    _dpi   string
    // 语言类型 zh_CN en_US
    _lang   string
}

// 初始化AlibabaSecurityJaqResourceFetchAPIRequest对象
func NewAlibabaSecurityJaqResourceFetchRequest() *AlibabaSecurityJaqResourceFetchAPIRequest{
    return &AlibabaSecurityJaqResourceFetchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.resource.fetch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceType Setter
// 设备类型可能值有：android ios wp
func (r *AlibabaSecurityJaqResourceFetchAPIRequest) SetDeviceType(_deviceType string) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetDeviceType() string {
    return r._deviceType
}
// Dpi Setter
// 分辨率
func (r *AlibabaSecurityJaqResourceFetchAPIRequest) SetDpi(_dpi string) error {
    r._dpi = _dpi
    r.Set("dpi", _dpi)
    return nil
}

// Dpi Getter
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetDpi() string {
    return r._dpi
}
// Lang Setter
// 语言类型 zh_CN en_US
func (r *AlibabaSecurityJaqResourceFetchAPIRequest) SetLang(_lang string) error {
    r._lang = _lang
    r.Set("lang", _lang)
    return nil
}

// Lang Getter
func (r AlibabaSecurityJaqResourceFetchAPIRequest) GetLang() string {
    return r._lang
}
