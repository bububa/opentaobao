package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备扩展信息 APIResponse
taobao.ailab.aicloud.top.device.extinfo.get

获取设备扩展信息
*/
type TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_device_extinfo_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoAilabAicloudTopDeviceExtinfoGetResult `json:"result,omitempty" xml:"