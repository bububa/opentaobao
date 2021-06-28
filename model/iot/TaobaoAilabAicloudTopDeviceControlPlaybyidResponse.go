package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过id播放歌曲 APIResponse
taobao.ailab.aicloud.top.device.control.playbyid

通过id播放歌曲
*/
type TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceControlPlaybyidResponse
}

type TaobaoAilabAicloudTopDeviceControlPlaybyidResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_control_playbyid_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
