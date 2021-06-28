package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取店铺关注链接 APIResponse
tmall.device.store.followurl.get

获取智能硬件上的关注店铺的URL
*/
type TmallDeviceStoreFollowurlGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallDeviceStoreFollowurlGetResponse `json:"tmall_device_store_followurl_get_response,omitempty"` 
    TmallDeviceStoreFollowurlGetResponse
}

/* model for simplify = false
type TmallDeviceStoreFollowurlGetResponse struct {

    // 二维码短链接地址
    
    ShortUrl   string `json:"short_url,omitempty"`
    

    // 二维码图片URL
    
    ShortImgUrl   string `json:"short_img_url,omitempty"`
    

}
*/

type TmallDeviceStoreFollowurlGetResponse struct {

    // 二维码短链接地址
    ShortUrl   string `json:"short_url,omitempty"`

    // 二维码图片URL
    ShortImgUrl   string `json:"short_img_url,omitempty"`

}
