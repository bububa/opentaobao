package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取二维码短链接 API请求
tmall.device.shortchainurl.get

获取二维码短链接
*/
type TmallDeviceShortchainurlGetRequest struct {
    model.Params
    // 是否需要长期二维码，默认否
    longterm   bool
    // 需要生成短链接的url
    url   string
    // 设备DeviceCode
    deviceCode   string
    // 商户中心门店ID
    storeId   int64
    // 动作类型，支持自定义
    action   string
}

// 初始化TmallDeviceShortchainurlGetRequest对象
func NewTmallDeviceShortchainurlGetRequest() *TmallDeviceShortchainurlGetRequest{
    return &TmallDeviceShortchainurlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDeviceShortchainurlGetRequest) GetApiMethodName() string {
    return "tmall.device.shortchainurl.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallDeviceShortchainurlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Longterm Setter
// 是否需要长期二维码，默认否
func (r *TmallDeviceShortchainurlGetRequest) SetLongterm(longterm bool) error {
    r.longterm = longterm
    r.Set("longterm", longterm)
    return nil
}

// Longterm Getter
func (r TmallDeviceShortchainurlGetRequest) GetLongterm() bool {
    return r.longterm
}
// Url Setter
// 需要生成短链接的url
func (r *TmallDeviceShortchainurlGetRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

// Url Getter
func (r TmallDeviceShortchainurlGetRequest) GetUrl() string {
    return r.url
}
// DeviceCode Setter
// 设备DeviceCode
func (r *TmallDeviceShortchainurlGetRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

// DeviceCode Getter
func (r TmallDeviceShortchainurlGetRequest) GetDeviceCode() string {
    return r.deviceCode
}
// StoreId Setter
// 商户中心门店ID
func (r *TmallDeviceShortchainurlGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TmallDeviceShortchainurlGetRequest) GetStoreId() int64 {
    return r.storeId
}
// Action Setter
// 动作类型，支持自定义
func (r *TmallDeviceShortchainurlGetRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

// Action Getter
func (r TmallDeviceShortchainurlGetRequest) GetAction() string {
    return r.action
}
