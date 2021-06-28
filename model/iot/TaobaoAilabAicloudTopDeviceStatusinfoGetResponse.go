package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取设备状态信息 APIResponse
taobao.ailab.aicloud.top.device.statusinfo.get

获取设备状态信息
*/
type TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceStatusinfoGetResponse `json:"ailab_aicloud_top_device_statusinfo_get_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceStatusinfoGetResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceStatusinfoGetResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoAilabAicloudTopDeviceStatusinfoGetResult  *TaobaoAilabAicloudTopDeviceStatusinfoGetResult `json:"taobao_ailab_aicloud_top_device_statusinfo_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceStatusinfoGetResponse struct {

    // 接口返回model
    Result   *TaobaoAilabAicloudTopDeviceStatusinfoGetResult `json:"result,omitempty"`

}
