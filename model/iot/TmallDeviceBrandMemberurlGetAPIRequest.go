package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取智能硬件旗舰店入会码 API请求
tmall.device.brand.memberurl.get

获取旗舰店在智能硬件上的入会码
*/
type TmallDeviceBrandMemberurlGetAPIRequest struct {
    model.Params
    // 设备DeviceCode
    _deviceCode   string
    // 入会后的回调地址
    _callbackUrl   string
    // 是否使用长期链接
    _longterm   bool
    // 页面banner的图片，如果没有传入，会使用系统默认图
    _bannerImg   string
    // 是否同时关注天猫理想站
    _followRetailAccount   bool
}

// 初始化TmallDeviceBrandMemberurlGetAPIRequest对象
func NewTmallDeviceBrandMemberurlGetRequest() *TmallDeviceBrandMemberurlGetAPIRequest{
    return &TmallDeviceBrandMemberurlGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDeviceBrandMemberurlGetAPIRequest) GetApiMethodName() string {
    return "tmall.device.brand.memberurl.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallDeviceBrandMemberurlGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备DeviceCode
func (r *TmallDeviceBrandMemberurlGetAPIRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r TmallDeviceBrandMemberurlGetAPIRequest) GetDeviceCode() string {
    return r._deviceCode
}
// CallbackUrl Setter
// 入会后的回调地址
func (r *TmallDeviceBrandMemberurlGetAPIRequest) SetCallbackUrl(_callbackUrl string) error {
    r._callbackUrl = _callbackUrl
    r.Set("callback_url", _callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r TmallDeviceBrandMemberurlGetAPIRequest) GetCallbackUrl() string {
    return r._callbackUrl
}
// Longterm Setter
// 是否使用长期链接
func (r *TmallDeviceBrandMemberurlGetAPIRequest) SetLongterm(_longterm bool) error {
    r._longterm = _longterm
    r.Set("longterm", _longterm)
    return nil
}

// Longterm Getter
func (r TmallDeviceBrandMemberurlGetAPIRequest) GetLongterm() bool {
    return r._longterm
}
// BannerImg Setter
// 页面banner的图片，如果没有传入，会使用系统默认图
func (r *TmallDeviceBrandMemberurlGetAPIRequest) SetBannerImg(_bannerImg string) error {
    r._bannerImg = _bannerImg
    r.Set("banner_img", _bannerImg)
    return nil
}

// BannerImg Getter
func (r TmallDeviceBrandMemberurlGetAPIRequest) GetBannerImg() string {
    return r._bannerImg
}
// FollowRetailAccount Setter
// 是否同时关注天猫理想站
func (r *TmallDeviceBrandMemberurlGetAPIRequest) SetFollowRetailAccount(_followRetailAccount bool) error {
    r._followRetailAccount = _followRetailAccount
    r.Set("follow_retail_account", _followRetailAccount)
    return nil
}

// FollowRetailAccount Getter
func (r TmallDeviceBrandMemberurlGetAPIRequest) GetFollowRetailAccount() bool {
    return r._followRetailAccount
}
