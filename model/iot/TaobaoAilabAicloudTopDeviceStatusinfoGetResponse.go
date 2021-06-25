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
    Response *TaobaoAilabAicloudTopDeviceStatusinfoGetResponse `json:"taobao_ailab_aicloud_top_device_statusinfo_get_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceStatusinfoGetResponse struct {

    // 接口返回model
    Result   *TaobaoAilabAicloudTopDeviceStatusinfoGetResult `json:"result,omitempty"`

}
