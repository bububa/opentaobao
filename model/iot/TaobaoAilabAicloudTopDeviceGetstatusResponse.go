package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取设备状态 APIResponse
taobao.ailab.aicloud.top.device.getstatus

获取设备状态
*/
type TaobaoAilabAicloudTopDeviceGetstatusAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopDeviceGetstatusResponse `json:"taobao_ailab_aicloud_top_device_getstatus_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceGetstatusResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
