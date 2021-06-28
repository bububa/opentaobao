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
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_device_brand_memberurl_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 二维码图片URL
    
    ShortImgUrl   string `json:"short_img_url,omitempty" xml:"