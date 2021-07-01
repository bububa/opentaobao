package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备播放暂停 API返回值 
taobao.ailab.aicloud.top.device.control.pauseandresume

设备播放暂停
*/
type TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponseModel
}

// 设备播放暂停 成功返回结果
type TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponseModel struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_control_pauseandresume_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
