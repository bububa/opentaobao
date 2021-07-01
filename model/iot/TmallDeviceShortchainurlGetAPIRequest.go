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
type TmallDeviceShortchainurlGetAPIRequest struct {
    model.Params
    // 是否需要长期二维码，默认否
    _longterm   bool
    // 需要生成短链接的url
    _url   string
    // 设备DeviceCode
    _deviceCode   string
    // 商户中心门店ID
    _storeId   int64
    // 动作类型，支持自定义
    _action   string
}

// 初始化TmallDeviceShortchainurlGetAPIRequest对象
func NewTmallDeviceShortchainurlGetRequest() *TmallDeviceShortchainurlGetAPIRequest{
    return &TmallDeviceShortchainurlGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDeviceShortchainurlGetAPIRequest) GetApiMethodName() string {
    return "tmall.device.shortchainurl.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallDeviceShortchainurlGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Longterm Setter
// 是否需要长期二维码，默认否
func (r *TmallDeviceShortchainurlGetAPIRequest) SetLongterm(_longterm bool) error {
    r._longterm = _longterm
    r.Set("longterm", _longterm)
    return nil
}

// Longterm Getter
func (r TmallDeviceShortchainurlGetAPIRequest) GetLongterm() bool {
    return r._longterm
}
// Url Setter
// 需要生成短链接的url
func (r *TmallDeviceShortchainurlGetAPIRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r TmallDeviceShortchainurlGetAPIRequest) GetUrl() string {
    return r._url
}
// DeviceCode Setter
// 设备DeviceCode
func (r *TmallDeviceShortchainurlGetAPIRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r TmallDeviceShortchainurlGetAPIRequest) GetDeviceCode() string {
    return r._deviceCode
}
// StoreId Setter
// 商户中心门店ID
func (r *TmallDeviceShortchainurlGetAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TmallDeviceShortchainurlGetAPIRequest) GetStoreId() int64 {
    return r._storeId
}
// Action Setter
// 动作类型，支持自定义
func (r *TmallDeviceShortchainurlGetAPIRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r TmallDeviceShortchainurlGetAPIRequest) GetAction() string {
    return r._action
}
