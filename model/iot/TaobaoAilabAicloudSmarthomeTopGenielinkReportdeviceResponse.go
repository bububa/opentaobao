package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
零配方案上报设备 APIResponse
taobao.ailab.aicloud.smarthome.top.genielink.reportdevice

零配方案中设备联网成功之后上报设备
*/
type TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResponse `json:"ailab_aicloud_smarthome_top_genielink_reportdevice_response,omitempty"` 
    TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResponse
}

/* model for simplify = false
type TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResponse struct {

    // result
    
    Result  *struct {
        TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult  *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult `json:"taobao_ailab_aicloud_smarthome_top_genielink_reportdevice_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResponse struct {

    // result
    Result   *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult `json:"result,omitempty"`

}
