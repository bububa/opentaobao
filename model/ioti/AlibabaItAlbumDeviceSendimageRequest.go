package ioti

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
相框设备厂测刷图接口 API请求
alibaba.it.album.device.sendimage

提供传入电子相框设备mac，mac需属于厂测白名单设备，将设备刷新为系统默认的厂测图片
*/
type AlibabaItAlbumDeviceSendimageRequest struct {
    model.Params
    // 下发图片mac地址
    mac   string
}

// 初始化AlibabaItAlbumDeviceSendimageRequest对象
func NewAlibabaItAlbumDeviceSendimageRequest() *AlibabaItAlbumDeviceSendimageRequest{
    return &AlibabaItAlbumDeviceSendimageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItAlbumDeviceSendimageRequest) GetApiMethodName() string {
    return "alibaba.it.album.device.sendimage"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItAlbumDeviceSendimageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mac Setter
// 下发图片mac地址
func (r *AlibabaItAlbumDeviceSendimageRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

// Mac Getter
func (r AlibabaItAlbumDeviceSendimageRequest) GetMac() string {
    return r.mac
}
