package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取二维码短链接 API返回值 
tmall.device.shortchainurl.get

获取二维码短链接
*/
type TmallDeviceShortchainurlGetAPIResponse struct {
    model.CommonResponse
    TmallDeviceShortchainurlGetAPIResponseModel
}

// 获取二维码短链接 成功返回结果
type TmallDeviceShortchainurlGetAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_device_shortchainurl_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 二维码图片地址
    ShortImgUrl   string `json:"short_img_url,omitempty" xml:"short_img_url,omitempty"`
    // 短链接url
    ShortUrl   string `json:"short_url,omitempty" xml:"short_url,omitempty"`
}
