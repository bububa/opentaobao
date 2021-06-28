package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取智能硬件旗舰店入会码 APIResponse
tmall.device.brand.memberurl.get

获取旗舰店在智能硬件上的入会码
*/
type TmallDeviceBrandMemberurlGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallDeviceBrandMemberurlGetResponse `json:"tmall_device_brand_memberurl_get_response,omitempty"` 
    TmallDeviceBrandMemberurlGetResponse
}

/* model for simplify = false
type TmallDeviceBrandMemberurlGetResponse struct {

    // 二维码图片URL
    
    ShortImgUrl   string `json:"short_img_url,omitempty"`
    

    // 二维码短链接地址
    
    ShortUrl   string `json:"short_url,omitempty"`
    

}
*/

type TmallDeviceBrandMemberurlGetResponse struct {

    // 二维码图片URL
    ShortImgUrl   string `json:"short_img_url,omitempty"`

    // 二维码短链接地址
    ShortUrl   string `json:"short_url,omitempty"`

}
