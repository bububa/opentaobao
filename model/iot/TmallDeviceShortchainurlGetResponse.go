package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取二维码短链接 APIResponse
tmall.device.shortchainurl.get

获取二维码短链接
*/
type TmallDeviceShortchainurlGetAPIResponse struct {
    model.CommonResponse
    TmallDeviceShortchainurlGetResponse
}

type TmallDeviceShortchainurlGetResponse struct {
    XMLName xml.Name `xml:"tmall_device_shortchainurl_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 二维码图片地址
    
    ShortImgUrl   string `json:"short_img_url,omitempty" xml:"short_img_url,omitempty"`

    
    // 短链接url
    
    ShortUrl   string `json:"short_url,omitempty" xml:"short_url,omitempty"`

    
}
