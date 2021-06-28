package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备状态 APIResponse
taobao.ailab.aicloud.top.device.getstatus

获取设备状态
*/
type TaobaoAilabAicloudTopDeviceGetstatusAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceGetstatusResponse
}

type TaobaoAilabAicloudTopDeviceGetstatusResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_getstatus_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
