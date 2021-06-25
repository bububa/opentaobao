package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取二维码短链接 APIRequest
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

func NewTmallDeviceShortchainurlGetRequest() *TmallDeviceShortchainurlGetRequest{
    return &TmallDeviceShortchainurlGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallDeviceShortchainurlGetRequest) GetApiMethodName() string {
    return "tmall.device.shortchainurl.get"
}

func (r TmallDeviceShortchainurlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallDeviceShortchainurlGetRequest) SetLongterm(longterm bool) error {
    r.longterm = longterm
    r.Set("longterm", longterm)
    return nil
}

func (r TmallDeviceShortchainurlGetRequest) GetLongterm() bool {
    return r.longterm
}

func (r *TmallDeviceShortchainurlGetRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r TmallDeviceShortchainurlGetRequest) GetUrl() string {
    return r.url
}

func (r *TmallDeviceShortchainurlGetRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r TmallDeviceShortchainurlGetRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *TmallDeviceShortchainurlGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TmallDeviceShortchainurlGetRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TmallDeviceShortchainurlGetRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

func (r TmallDeviceShortchainurlGetRequest) GetAction() string {
    return r.action
}

