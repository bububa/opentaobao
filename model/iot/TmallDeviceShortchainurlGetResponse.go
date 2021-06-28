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
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_device_shortchainurl_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 二维码图片地址
    
    ShortImgUrl   string `json:"short_img_url,omitempty" xml:"