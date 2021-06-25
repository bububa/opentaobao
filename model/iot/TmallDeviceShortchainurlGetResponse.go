package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取二维码短链接 APIResponse
tmall.device.shortchainurl.get

获取二维码短链接
*/
type TmallDeviceShortchainurlGetAPIResponse struct {
    model.CommonResponse
    Response *TmallDeviceShortchainurlGetResponse `json:"tmall_device_shortchainurl_get_response,omitempty"`
}

type TmallDeviceShortchainurlGetResponse struct {

    // 二维码图片地址
    ShortImgUrl   string `json:"short_img_url,omitempty"`

    // 短链接url
    ShortUrl   string `json:"short_url,omitempty"`

}