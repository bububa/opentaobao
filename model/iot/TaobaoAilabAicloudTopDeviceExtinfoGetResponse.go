package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取设备扩展信息 APIResponse
taobao.ailab.aicloud.top.device.extinfo.get

获取设备扩展信息
*/
type TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopDeviceExtinfoGetResponse `json:"taobao_ailab_aicloud_top_device_extinfo_get_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceExtinfoGetResponse struct {

    // 接口返回model
    Result   *TaobaoAilabAicloudTopDeviceExtinfoGetResult `json:"result,omitempty"`

}
