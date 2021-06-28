package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取智能硬件旗舰店入会码 APIResponse
tmall.device.brand.memberurl.get

获取旗舰店在智能硬件上的入会码
*/
type TmallDeviceBrandMemberurlGetAPIResponse struct {
    model.CommonResponse
    TmallDeviceBrandMemberurlGetResponse
}

type TmallDeviceBrandMemberurlGetResponse struct {
    XMLName xml.Name `xml:"tmall_device_brand_memberurl_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 二维码图片URL
    
    ShortImgUrl   string `json:"short_img_url,omitempty" xml:"short_img_url,omitempty"`

    
    // 二维码短链接地址
    
    ShortUrl   string `json:"short_url,omitempty" xml:"short_url,omitempty"`

    
}
