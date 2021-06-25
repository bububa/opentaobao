package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取店铺关注链接 APIRequest
tmall.device.store.followurl.get

获取智能硬件上的关注店铺的URL
*/
type TmallDeviceStoreFollowurlGetRequest struct {
    model.Params

    // 设备DeviceCode
    deviceCode   string 

    // 关注完成后的回调地址,需要是EWS地址
    callbackUrl   string 

    // 是否同时关注天猫理想站
    followRetailAccount   bool 

    // 是否使用长期链接
    longterm   bool 

    // 页面banner的图片，如果没有传入，会使用系统默认图
    bannerImg   string 

}

func NewTmallDeviceStoreFollowurlGetRequest() *TmallDeviceStoreFollowurlGetRequest{
    return &TmallDeviceStoreFollowurlGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallDeviceStoreFollowurlGetRequest) GetApiMethodName() string {
    return "tmall.device.store.followurl.get"
}

func (r TmallDeviceStoreFollowurlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallDeviceStoreFollowurlGetRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r TmallDeviceStoreFollowurlGetRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *TmallDeviceStoreFollowurlGetRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

func (r TmallDeviceStoreFollowurlGetRequest) GetCallbackUrl() string {
    return r.callbackUrl
}

func (r *TmallDeviceStoreFollowurlGetRequest) SetFollowRetailAccount(followRetailAccount bool) error {
    r.followRetailAccount = followRetailAccount
    r.Set("follow_retail_account", followRetailAccount)
    return nil
}

func (r TmallDeviceStoreFollowurlGetRequest) GetFollowRetailAccount() bool {
    return r.followRetailAccount
}

func (r *TmallDeviceStoreFollowurlGetRequest) SetLongterm(longterm bool) error {
    r.longterm = longterm
    r.Set("longterm", longterm)
    return nil
}

func (r TmallDeviceStoreFollowurlGetRequest) GetLongterm() bool {
    return r.longterm
}

func (r *TmallDeviceStoreFollowurlGetRequest) SetBannerImg(bannerImg string) error {
    r.bannerImg = bannerImg
    r.Set("banner_img", bannerImg)
    return nil
}

func (r TmallDeviceStoreFollowurlGetRequest) GetBannerImg() string {
    return r.bannerImg
}

