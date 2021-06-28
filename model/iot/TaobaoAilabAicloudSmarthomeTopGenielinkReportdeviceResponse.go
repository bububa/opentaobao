package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零配方案上报设备 APIResponse
taobao.ailab.aicloud.smarthome.top.genielink.reportdevice

零配方案中设备联网成功之后上报设备
*/
type TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_smarthome_top_genielink_reportdevice_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult `json:"result,omitempty" xml:"