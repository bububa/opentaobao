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
type TmallDeviceStoreFollowurlGetRequest struct {
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

// 初始化TmallDeviceStoreFollowurlGetRequest对象
func NewTmallDeviceStoreFollowurlGetRequest() *TmallDeviceStoreFollowurlGetRequest{
    return &TmallDeviceStoreFollowurlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDeviceStoreFollowurlGetRequest) GetApiMethodName() string {
    return "tmall.device.store.followurl.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallDeviceStoreFollowurlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备DeviceCode
func (r *TmallDeviceStoreFollowurlGetRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r TmallDeviceStoreFollowurlGetRequest) GetDeviceCode() string {
    return r._deviceCode
}
// CallbackUrl Setter
// 关注完成后的回调地址,需要是EWS地址
func (r *TmallDeviceStoreFollowurlGetRequest) SetCallbackUrl(_callbackUrl string) error {
    r._callbackUrl = _callbackUrl
    r.Set("callback_url", _callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r TmallDeviceStoreFollowurlGetRequest) GetCallbackUrl() string {
    return r._callbackUrl
}
// FollowRetailAccount Setter
// 是否同时关注天猫理想站
func (r *TmallDeviceStoreFollowurlGetRequest) SetFollowRetailAccount(_followRetailAccount bool) error {
    r._followRetailAccount = _followRetailAccount
    r.Set("follow_retail_account", _followRetailAccount)
    return nil
}

// FollowRetailAccount Getter
func (r TmallDeviceStoreFollowurlGetRequest) GetFollowRetailAccount() bool {
    return r._followRetailAccount
}
// Longterm Setter
// 是否使用长期链接
func (r *TmallDeviceStoreFollowurlGetRequest) SetLongterm(_longterm bool) error {
    r._longterm = _longterm
    r.Set("longterm", _longterm)
    return nil
}

// Longterm Getter
func (r TmallDeviceStoreFollowurlGetRequest) GetLongterm() bool {
    return r._longterm
}
// BannerImg Setter
// 页面banner的图片，如果没有传入，会使用系统默认图
func (r *TmallDeviceStoreFollowurlGetRequest) SetBannerImg(_bannerImg string) error {
    r._bannerImg = _bannerImg
    r.Set("banner_img", _bannerImg)
    return nil
}

// BannerImg Getter
func (r TmallDeviceStoreFollowurlGetRequest) GetBannerImg() string {
    return r._bannerImg
}
