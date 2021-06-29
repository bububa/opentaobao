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
type TmallDeviceBrandMemberurlGetRequest struct {
    model.Params
    // 设备DeviceCode
    deviceCode   string
    // 入会后的回调地址
    callbackUrl   string
    // 是否使用长期链接
    longterm   bool
    // 页面banner的图片，如果没有传入，会使用系统默认图
    bannerImg   string
    // 是否同时关注天猫理想站
    followRetailAccount   bool
}

// 初始化TmallDeviceBrandMemberurlGetRequest对象
func NewTmallDeviceBrandMemberurlGetRequest() *TmallDeviceBrandMemberurlGetRequest{
    return &TmallDeviceBrandMemberurlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDeviceBrandMemberurlGetRequest) GetApiMethodName() string {
    return "tmall.device.brand.memberurl.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallDeviceBrandMemberurlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 设备DeviceCode
func (r *TmallDeviceBrandMemberurlGetRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

// DeviceCode Getter
func (r TmallDeviceBrandMemberurlGetRequest) GetDeviceCode() string {
    return r.deviceCode
}
// CallbackUrl Setter
// 入会后的回调地址
func (r *TmallDeviceBrandMemberurlGetRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r TmallDeviceBrandMemberurlGetRequest) GetCallbackUrl() string {
    return r.callbackUrl
}
// Longterm Setter
// 是否使用长期链接
func (r *TmallDeviceBrandMemberurlGetRequest) SetLongterm(longterm bool) error {
    r.longterm = longterm
    r.Set("longterm", longterm)
    return nil
}

// Longterm Getter
func (r TmallDeviceBrandMemberurlGetRequest) GetLongterm() bool {
    return r.longterm
}
// BannerImg Setter
// 页面banner的图片，如果没有传入，会使用系统默认图
func (r *TmallDeviceBrandMemberurlGetRequest) SetBannerImg(bannerImg string) error {
    r.bannerImg = bannerImg
    r.Set("banner_img", bannerImg)
    return nil
}

// BannerImg Getter
func (r TmallDeviceBrandMemberurlGetRequest) GetBannerImg() string {
    return r.bannerImg
}
// FollowRetailAccount Setter
// 是否同时关注天猫理想站
func (r *TmallDeviceBrandMemberurlGetRequest) SetFollowRetailAccount(followRetailAccount bool) error {
    r.followRetailAccount = followRetailAccount
    r.Set("follow_retail_account", followRetailAccount)
    return nil
}

// FollowRetailAccount Getter
func (r TmallDeviceBrandMemberurlGetRequest) GetFollowRetailAccount() bool {
    return r.followRetailAccount
}
