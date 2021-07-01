package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取店铺关注链接 API请求
tmall.device.store.followurl.get

获取智能硬件上的关注店铺的URL
*/
type TmallDeviceStoreFollowurlGetAPIRequest struct {
    model.Params
    // 设备DeviceCode
    _deviceCode   string
    // 关注完成后的回调地址,需要是EWS地址
    _callbackUrl   string
    // 是否同时关注天猫理想站
    _followRetailAccount   bool
    // 是否使用长期链接
    _longterm   bool
    // 页面banner的图片，如果没有传入，会使用系统默认图
    _bannerImg   string
}

// 初始化TmallDeviceStoreFollowurlGetAPIRequest对象
func NewTmallDeviceStoreFollowurlGetRequest() *TmallDeviceStoreFollowurlGetAPIRequest{
    return &TmallDeviceStoreFollowurlGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDeviceStoreFollowurlGetAPIRequest) GetApiMethodName() string {
    return "tmall.device.store.followurl.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallDeviceStoreFollowurlGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备DeviceCode
func (r *TmallDeviceStoreFollowurlGetAPIRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r TmallDeviceStoreFollowurlGetAPIRequest) GetDeviceCode() string {
    return r._deviceCode
}
// CallbackUrl Setter
// 关注完成后的回调地址,需要是EWS地址
func (r *TmallDeviceStoreFollowurlGetAPIRequest) SetCallbackUrl(_callbackUrl string) error {
    r._callbackUrl = _callbackUrl
    r.Set("callback_url", _callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r TmallDeviceStoreFollowurlGetAPIRequest) GetCallbackUrl() string {
    return r._callbackUrl
}
// FollowRetailAccount Setter
// 是否同时关注天猫理想站
func (r *TmallDeviceStoreFollowurlGetAPIRequest) SetFollowRetailAccount(_followRetailAccount bool) error {
    r._followRetailAccount = _followRetailAccount
    r.Set("follow_retail_account", _followRetailAccount)
    return nil
}

// FollowRetailAccount Getter
func (r TmallDeviceStoreFollowurlGetAPIRequest) GetFollowRetailAccount() bool {
    return r._followRetailAccount
}
// Longterm Setter
// 是否使用长期链接
func (r *TmallDeviceStoreFollowurlGetAPIRequest) SetLongterm(_longterm bool) error {
    r._longterm = _longterm
    r.Set("longterm", _longterm)
    return nil
}

// Longterm Getter
func (r TmallDeviceStoreFollowurlGetAPIRequest) GetLongterm() bool {
    return r._longterm
}
// BannerImg Setter
// 页面banner的图片，如果没有传入，会使用系统默认图
func (r *TmallDeviceStoreFollowurlGetAPIRequest) SetBannerImg(_bannerImg string) error {
    r._bannerImg = _bannerImg
    r.Set("banner_img", _bannerImg)
    return nil
}

// BannerImg Getter
func (r TmallDeviceStoreFollowurlGetAPIRequest) GetBannerImg() string {
    return r._bannerImg
}
