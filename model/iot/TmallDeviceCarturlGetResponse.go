package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加商品到购物车 APIResponse
tmall.device.carturl.get

获取二维码，支持添加商品到购物车
*/
type TmallDeviceCarturlGetAPIResponse struct {
    model.CommonResponse
    Response *TmallDeviceCarturlGetResponse `json:"tmall_device_carturl_get_response,omitempty"`
}

type TmallDeviceCarturlGetResponse struct {

    // 二维码短链接地址
    ShortUrl   string `json:"short_url,omitempty"`

    // 二维码图片URL
    ShortImgUrl   string `json:"short_img_url,omitempty"`

}
