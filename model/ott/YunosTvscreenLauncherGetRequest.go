package ott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
一体机桌面 API请求
yunos.tvscreen.launcher.get

LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务
*/
type YunosTvscreenLauncherGetRequest struct {
    model.Params
    // 设备属性
    property   string
    // IP来源
    ip   string
}

// 初始化YunosTvscreenLauncherGetRequest对象
func NewYunosTvscreenLauncherGetRequest() *YunosTvscreenLauncherGetRequest{
    return &YunosTvscreenLauncherGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvscreenLauncherGetRequest) GetApiMethodName() string {
    return "yunos.tvscreen.launcher.get"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvscreenLauncherGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Property Setter
// 设备属性
func (r *YunosTvscreenLauncherGetRequest) SetProperty(property string) error {
    r.property = property
    r.Set("property", property)
    return nil
}

// Property Getter
func (r YunosTvscreenLauncherGetRequest) GetProperty() string {
    return r.property
}
// Ip Setter
// IP来源
func (r *YunosTvscreenLauncherGetRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r YunosTvscreenLauncherGetRequest) GetIp() string {
    return r.ip
}
