package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详细信息 APIResponse
taobao.ailab.aicloud.top.device.detailinfo.get

获取设备详细信息
*/
type TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopDeviceDetailinfoGetResponse `json:"taobao_ailab_aicloud_top_device_detailinfo_get_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceDetailinfoGetResponse struct {

    // 接口返回model
    Result   *TaobaoAilabAicloudTopDeviceDetailinfoGetResult `json:"result,omitempty"`

}
